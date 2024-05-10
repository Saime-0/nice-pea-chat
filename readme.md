1. Дизайн системы

1.1 Функционал



1.1.1 Клиент - мобильное приложение (android)

a.k.a mobile app

Экран выбора сервера и ввода логина  -

Открывается при первом входе в приложение

Выбранный сервер сохраняется в кэш

Компоненты:

Инпут для адреса

История уникальных адресов

Инпут логина (появляется после удачного соединения с сервером)

Кнопка "далее"

После ввода и подтверждения сервера, появляется инпут ввода логина, если адрес будет отредактирован, то логин пропадет

После ввода логина и нажатия далее, откроется экран списка чатов

Главный экран со списком чатов -

Компоненты:

Список чатов

По нажатию на чат, откроется его экран сообщений

Отсортирован по времени сообщений, последний сверх, а также по закрепленным, закрепленные сверху

Каждый элемент состоит из:

Названия

Индикатора количества новых сообщений

Автора (если есть) и текста сообщения

Кнопки "закрепить"

Кнопка "посмотреть профиль"

Открывает диалог профиля пользователя

Диалог профиля пользователя -

Компоненты:

Идентификатор

Юзернейм

Логин

Кнопка выхода

Выбрасывает на экран логина

Экран сообщений чата -

Компоненты:

Инпут бар сообщения

Кнопка "вернуться"

Список сообщений:

Каждый элемент состоит из:

Имени отправителя

Части сообщения на которое был сделан ответ (необязательно)

Текста сообщения

Даты отправки

Даты редактирования/удаления (необязательно)

По нажатию открывается меню

Компоненты:

Кнопка скопировать сообщение

Кнопка ответить на сообщение

Кнопки удалить и изменить сообщение (необязательно)

По нажатию, появится диалог с подтверждением

Приложение -

Если запрос прошел с ошибкой (не 500), то текст из ответа надо вывести на экрана

Если сервер вернет ошибку авторизации, приложение должно очистить кэш (кроме логина и адреса) и перейти на экран логина

1.1.2 Чат

a.k.a conversatin/беседа

Структура чата:

Идентификатор

идентификатор создателя

Название чата

Список участников

Профиль чата (редактируемые поля):

Название чата

Приглашение в чат (invite):

Приглашать в чат могут только уполномоченные пользователи

Новый участник чата не получает уведомлений, до тех пор пока не прочтет хотя бы одно сообщение

Приглашение пользователя в чат, принудительно добавляет его, т.е отказать от приглашения невозможно

Чат -

В чат можно попасть по приглашению, либо создав новый чат, т.е все чаты являются приватными

Создатель чата не может покинуть свой чат

Количество чатов у пользователя не ограничено

Количество созданных чатов ограничено, лимит равен 100

Управление чатом и его сообщениями опирается на permissions

У пользователя должна быть возможность закреплять свои чаты вверху списка чатов в которых он состоит

1.1.3 Права участников

a.k.a разрешения



Набор прав (permission set) -

Набор прав можно выдавать участнику чата, но одновременно не может быть больше 1 набора прав

Выдать права можно только имея на это права

Permissions set - составляется в system configuration, т.е пользователи его меня не могут, только использовать заготовленные варианты

Права (permissions) -

Права дают доступ на какое-либо действие

Некоторые требуют передачи параметра target, в этом случае помимо наличия права, будет выполнена проверка - target (участник) должен обладать уровнем прав ниже, чтобы выполняемое действие прошло успешно

В качестве параметра target разрешено передавать только участников, чей уровень прав ниже своего

Права объединенные в группу - набор прав (permission set)

Permissions set - создается в system configuration

Пример прав:

MsgSend

MsgOwnEdit

MsgOwnDelete

MsgDelete(target)

ChatProfileEdit

MemberDelete(target)

MemberAdd

GivePermissionSet(target)

Пример permission set:

Member(0): MsgSend, MsgOwnEdit, MsgOwnDelete

Moder(100): Member + MsgOtherDelete, MemberAdd

Moder2(200): Moderator + MemberDelete, ChatProfileEdit

Admin(300): Moder2 + GivePermissionSet

1.1.4 Системные настройки

a.k.a System configuration



1.1.5 Участник чата

a.k.a member

Структура участника:

Идентификатор

Идентификатор пользователя

Идентификатор чата

permission set

Участник -

Факт отношения пользователя к определенному чату

1.1.6 Сообщение

a.k.a message

Типы событий, в результате которых приходит анонимное сообщение в чат:

Изменение профиля чата

Создание нового участника

Удаление участника

Самостоятельный выход участника

Создание чата

Структура сообщения:

идентификатор

идентификатор чата

идентификатор пользователя, т.е автора сообщения (необязательно)

идентификатор сообщения, на которое это сообщение отвечает

текст (отсутствует если оно было удалено)

дата создания

дата редактирования (необязательно)

дата удаления (необязательно)

список пользователей, прочитавших сообщение (исключая создателя)

Сообщение (message) -

Может быть создано пользователем, являющимися участником чата, либо в результате события в чате

Текст может быть множество раз отредактирован автором, если не сообщение является удаленным

Может быть единожды удалено автором

На текст сообщения есть ограничение в 4096 символов

На существующее (даже если удалено) сообщение можно отвечать другим сообщением

Полная история сообщений в чате доступна каждому участнику

1.1.7 Push-уведомление

a.k.a notification/push-notification

Структура пуша:

Название чата

Текст сообщения (необязательно)

Текст события (необязательно)

В каких случаях пользователь получает пуш:

В любой чат в котором он состоит пришло новое сообщение

Приглашение пользователя (только тому кого пригласили)

Удаление пользователя (только тому кого удалили)

Когда пользователь не получает уведомления о новых сообщениях:

Если его только пригласили в чат (является новым участником), и он ни одного сообщения не прочел

Если пользователь уже находится в этом чате, т.е он открыт на клиенте

Если этот чат находится в муте у пользователя

Уведомления -

Уведомления создают push-уведомления на клиенте, даже если сейчас он закрыт

1.1.8 Пользователь

a.k.a user



Структура пользователя:

Идентификатор

Юзернейм

Логин

Профиль пользователя (редактируемые поля):

Юзернейм

Пользователь -

Пользователя можно создать пройдя регистрацию

"Вход в профиль" осуществляется по логину

Профиль можно редактировать

1.2 Интерфейс ui/ux

1.2.1 Заголовок темы

1.3 Техническая часть



API для приложения

1.3 Авторизация

1.3 Аутентификация

1.3 Ответ с ошибкой

1.3 Сообщение

Модель

Message:
  id: i32
  chat: i32
  user: i32?
  reply: i32?
  text: str
  date: i64
  edit_date: i64?
  delete_date: i64?
  readers: []i32



Пагинация сообщений в чате

Для выборки сообщений в определенном чате требуется передать startId и endId



1.3 Участник

Модель

Member:
  id: i32
  chat: i32
  user: i32
  permission_set: PermissionSet



1.3 Пользователь

Модель

User:
  id: i32
  username: str
  login: str



1.3 Разрешения

Модель

PermissionSet:
  id: i32
  lvl: u16
  name: str
  permissions: []Permission
 
Permission:
  name: str
  type: PermissionType
 
PermissionType:
  MsgSend
  MsgOwnEdit
  MsgOwnDelete
  MsgDelete
  ChatProfileEdit
  MemberDelete
  MemberAdd
  GivePermissionSet



1.3 Чат

Модель

Chat:
  id: i32
  owner: i32
  name: str







2. Сбор пожертвований требований

2.1 Технические вопросы

На чем основана аутентификация? jwt или другие варианты? Если jwt то какая нагрузка у токенов будет?

Один инстанс сервера будет использоваться для всех чатов, т.е. 1:n либо 1:1, если таковые будут.

Будет ли поддержка горизонтального масштабирования? Если да, то что будет шлюзом?

Какой api будет между сервером, http, grpc, gql?

Будет ли использоваться долгоживущее соединение - websocket, long polling, etc?

Сервер будет stateless или statefull?

Будет ли сервер писать логи, как их можно будет читать?

Что будет использоваться для хранения данных - сообщения, данные чатов, пользователей?

Сообщения

Как будет происходить пагинация при запросе сообщений?

2.2 Анализ требуемого функционала

Какой клиент будет у "Cute-chat", web, mobile, desktop?

"Cute-chat" это о одном чате, т.е. все сообщения будут сваливаться в одну кучу или будет возможность отправлять сообщение в разные чаты?

Несколько чатов

Если чатов несколько, то вероятно их кто то должен создавать, кто это будет делать, или у кого такая возможность будет?

Как пользователи будут попадать в чаты? По приглашению, присоединяться самостоятельно, присоединяться принудительно?

Ограничения на вход будут? Т.е. приватые/публичные чаты?

Если возможность создавать чаты будет у пользователя, то вероятно должны быть ограничения на количество созданных чатов?

Как выглядит профиль чата? В профиле будут название, картинка, описание, список участников?

Будут ли у создателя чата особые права? Например он может изменять профиль чата, удалять чат, добавлять\исключать участников, модерировать?

Может ли создатель чата выходить из чата? Что будет если создатель выйдет из чата?

Будет возможность передавать права создателя?

Возможность удалить чат будет? Что произойдет с участниками? Если в этот момент пользователь перейдет по приглашению в чат, что он увидит?

Есть ли возможность просматривать список участников чата? Кому эта возможность доступна?

В чате есть ограничения по количеству участников?

Можно ли пересылать сообщения между чатами?

Сообщения

Что в чат можно отправлять пользователям? Медиа/стикеры/ссылки/текст/файлы?

Сообщение может содержать несколько типов контента? Будет ли применяться разбиение на несколько сообщений при превышении лимита медиа или символов в тексте?

Есть ли ограничения на отправку сообщений?

Какие действия можно совершать с сообщениями? Удалять/изменять/пересылать/отвечать/копировать/отправлять повторно?

Как будет выглядеть пересланное сообщение если его оригинал будет изменен/удален?

Сохраняется ли история сообщений? Новые участники видят всю историю или только ее часть?

Какого типа сообщения в чате будут появляться? Пользовательские/системные?

Возможность отправлять анонимные сообщения в чат будет?

Возможность отправлять сообщения со спойлером будет?

Возможность отправлять сообщения с автоматическим удалением будет? Что произойдет если это сообщение будет переслано?

Сообщения будут обладать статусом "прочитано"?

Какой порядок сообщений в истории, последние помещаются снизу или сверху?

Как выглядит профиль сообщения? В профиле сообщения будут дата создания, флаг о том что его прочитали, дата изменения/удаления, автор сообщения, чат?

Поиск по сообщениям в конкретном чате будет? К поиску будет возможность добавлять фильтры, автор, даты, статус, тип (медиа\текст)?

В текстовых сообщениях можно будет упоминать пользователей?

Новое сообщение

Новое сообщение будет сопровождаться индикатором на чате, меткой "новое" в истории?

Новое сообщение будет рассылать push уведомления всем участникам? Возможность отключать push уведомления будет?

Пользователь

Откуда берется пользователь? Он создается исключительно для "cute-chat", т.е. проходит регистрацию или база пользователей уже существует?

Как выглядит профиль пользователя? Имя, почта, ид, фото, описание, список чатов? Какие данные пользователь может редактировать?

Будут ли существовать пользователи с дополнительными возможностями? Например с возможностью редактировать/просматривать/удалять то что обычные не могут

Клиент web/mobile/desktop

Как будет происходить связывание клиента и сервера, в клиенте будет статически прописан адрес или пользователям надо вручную указывать адрес сервера?

Будет ли ограничения при использовании клиента не последней версии либо отчающейс я от поддерживаемой  сервером?
