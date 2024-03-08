package services

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	e "vk2telegram/internal/errors"
	"vk2telegram/internal/model/messenger"
	"vk2telegram/internal/model/user"
	"vk2telegram/internal/storage"
)

const (
	TdUnknown = iota
	TdSubscription
	TdSearchGroup
	TdRegister
)

type TelegramAPI struct {
	bot *tgbotapi.BotAPI
	vk  *VKAPI
}

func NewTelegramAPI(token string, vk *VKAPI) *TelegramAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	return &TelegramAPI{bot: bot, vk: vk}
}

func (t *TelegramAPI) Send(post messenger.Post, user user.User) {
	var attachments []interface{}
	_, _ = t.bot.Send(tgbotapi.NewMessage(user.ChatID, post.Text))
	if len(post.Attachments) > 0 {
		for _, attachment := range post.Attachments {
			switch attachment.Type {
			case messenger.AttachmentTypePhoto:
				var priority int
				var s messenger.Size
				for _, size := range attachment.Photo.Sizes {
					if size.Type.Priority() > priority {
						priority = size.Type.Priority()
						s = size
					}
				}
				attachments = append(attachments, tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(s.Url)))
			case messenger.AttachmentTypeDocument:
				_, _ = t.bot.Send(tgbotapi.NewDocument(user.ChatID, tgbotapi.FileURL(attachment.Doc.Url)))
			}
		}
	}

	if len(attachments) > 0 {
		_, _ = t.bot.Send(tgbotapi.NewMediaGroup(user.ChatID, attachments))
	}
}

func (t *TelegramAPI) GetUpdates(db *storage.DatabaseStorage) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				var msg tgbotapi.MessageConfig
				u, _ := db.GetUser(update.Message.Chat.ID)
				if u == nil {
					text := fmt.Sprintf("Привет, %v %v!\n\nЭто бот, которые позволяет получать уведомления о новых постах из VK. Для этого необходимо указать токен пользователя VK и подписаться на сообщества. Всё, кроме получения токена, необходимо делать именно здесь.\n\nЕсли возникли какие-то вопросы или предложения, то можешь обратиться к разработчику https://t.me/karl_stein", update.Message.From.FirstName, update.Message.From.LastName)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
					msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Регистрация", strconv.Itoa(TdRegister))))
				} else {
					text := fmt.Sprintf("%v %v, вы уже зарегистрированы и можете полноценно пользоваться ботом", u.FirstName, u.LastName)
					keyboard := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("Подписки", strconv.Itoa(TdSubscription)),
						),
					)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
					msg.ReplyMarkup = keyboard
				}
				_, _ = t.bot.Send(msg)
			}
		}
		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case strconv.Itoa(TdRegister):
				var text string
				u, err := db.AddUser(update.CallbackQuery.Message.Chat.FirstName, update.CallbackQuery.Message.Chat.LastName, update.CallbackQuery.Message.Chat.ID)
				if err != nil {
					switch {
					case errors.Is(err, e.UserAlreadyRegistered):
						text = fmt.Sprintf("%v %v, вы уже зарегистрированы и можете полноценно пользоваться ботом", u.FirstName, u.LastName)
					default:
						text = "Произошла неизвестная ошибка"
						log.Fatalln(err.Error())
					}
				}
				if len(text) == 0 {
					text = "Вы успешно зарегистрированы"
				}
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Подписки", strconv.Itoa(TdSubscription)),
					),
				)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				msg.ReplyMarkup = keyboard
				_, _ = t.bot.Send(msg)
			case strconv.Itoa(TdSubscription):
				var text string

				u, err := db.GetUser(update.CallbackQuery.Message.Chat.ID)
				if err != nil || u == nil {
					log.Println(fmt.Sprintf("Пользователь не найден или произошла ошибка: %v", err.Error()))
					break
				}

				groups, err := db.GetGroups(u.ID)
				if err != nil {
					log.Println(err.Error())
					break
				}

				if len(groups) == 0 {
					text = "Вы не подписаны ни на одно сообщество"
				} else {
					text = "Сообщества, на которые вы подписаны:\n"
					for index, group := range groups {
						text += strconv.Itoa(index+1) + ". " + group.Name + " (https://vk.com/" + group.ScreenName + ")\n"
					}
				}

				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Поиск сообществ", strconv.Itoa(TdSearchGroup)),
					),
				)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				msg.ReplyMarkup = keyboard
				_, _ = t.bot.Send(msg)
			case strconv.Itoa(TdSearchGroup):
				text := "Пока что, поиск сообществ не работает"
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Подписки", strconv.Itoa(TdSubscription)),
						tgbotapi.NewInlineKeyboardButtonData("Поиск сообществ", strconv.Itoa(TdSearchGroup)),
					),
				)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				msg.ReplyMarkup = keyboard
				_, _ = t.bot.Send(msg)
			}
		}
	}
}
