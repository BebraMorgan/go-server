# go-server
## Пакет: main
# server
--
Package main содержит точку входа в приложение сервера для работы с пивом и
закусками. В функции main происходит инициализация базы данных, автоматическая
миграция моделей и запуск сервера.

---

## Пакет: controllers
# controllers
--
    import "."

Package controllers содержит HTTP-обработчики для работы с сущностью Beer.

Package controllers содержит HTTP-обработчики для работы с сущностью Snack.

## Usage

#### func  CreateSnack

```go
func CreateSnack(r *request.Request, params map[string]string) types.JsonResponse
```
CreateSnack создаёт новую запись закуски на основе JSON из запроса. Возвращает
ошибку, если входные данные некорректны или произошла ошибка базы данных.

#### func  DeleteBeer

```go
func DeleteBeer(r *request.Request, params map[string]string) types.JsonResponse
```
DeleteBeer удаляет запись пива по ID. Возвращает ошибку, если ID отсутствует,
некорректен или произошла ошибка при удалении.

#### func  DeleteSnack

```go
func DeleteSnack(r *request.Request, params map[string]string) types.JsonResponse
```
DeleteSnack удаляет запись закуски по ID. Возвращает ошибку, если ID
отсутствует, некорректен или произошла ошибка при удалении.

#### func  GetAllSnacks

```go
func GetAllSnacks(r *request.Request, params map[string]string) types.JsonResponse
```
GetAllSnacks возвращает список всех закусок из базы данных. В случае ошибки
возвращает соответствующее сообщение.

#### func  GetRandomBeer

```go
func GetRandomBeer(r *request.Request, params map[string]string) types.JsonResponse
```
GetRandomBeer возвращает случайное пиво из базы данных. Если пиво не найдено,
возвращает ошибку с соответствующим сообщением.

#### func  GetRandomSnack

```go
func GetRandomSnack(r *request.Request, params map[string]string) types.JsonResponse
```
GetRandomSnack возвращает случайную закуску из базы данных. Если закусок нет,
возвращает соответствующее сообщение об ошибке.

#### func  GetSnack

```go
func GetSnack(r *request.Request, params map[string]string) types.JsonResponse
```
GetSnack возвращает закуску по ID, переданному в параметрах маршрута. Возвращает
ошибку, если ID отсутствует, некорректен или запись не найдена.

#### func  ShowBeer

```go
func ShowBeer(r *request.Request, params map[string]string) types.JsonResponse
```
ShowBeer возвращает пиво по ID, переданному в параметрах маршрута. Возвращает
ошибку, если ID отсутствует, некорректен или запись не найдена.

#### func  StoreBeer

```go
func StoreBeer(r *request.Request, params map[string]string) types.JsonResponse
```
StoreBeer создаёт новую запись пива на основе JSON из запроса. Возвращает
ошибку, если входные данные некорректны или произошла ошибка базы данных.

#### func  UpdateBeer

```go
func UpdateBeer(r *request.Request, params map[string]string) types.JsonResponse
```
UpdateBeer обновляет существующую запись пива по ID. Возвращает ошибку, если ID
отсутствует, некорректен, запись не найдена или входные данные некорректны.

#### func  UpdateSnack

```go
func UpdateSnack(r *request.Request, params map[string]string) types.JsonResponse
```
UpdateSnack обновляет существующую запись закуски по ID. Возвращает ошибку, если
ID отсутствует, некорректен, запись не найдена или входные данные некорректны.

---

## Пакет: database
# database
--
    import "."

Package database отвечает за инициализацию и настройку подключения к базе
данных.

## Usage

```go
var DB *gorm.DB
```
DB — глобальный экземпляр подключения к базе данных.

#### func  Init

```go
func Init() error
```
Init инициализирует подключение к базе данных MySQL с заданными параметрами,
настраивает пул соединений и логирование. Возвращает ошибку в случае неудачи
подключения.

---

## Пакет: ./models
# models
--
    import "."

Package models содержит определения моделей данных и функции для работы с ними.

Package models содержит определения моделей данных и функции для работы с ними.

## Usage

#### func  CreateBeer

```go
func CreateBeer(db *gorm.DB, beer *Beer) error
```
CreateBeer сохраняет новую запись пива в базе данных.

#### func  CreateSnack

```go
func CreateSnack(db *gorm.DB, snack *Snack) error
```
CreateSnack сохраняет новую запись закуски в базе данных.

#### func  DeleteBeer

```go
func DeleteBeer(db *gorm.DB, id uint) error
```
DeleteBeer удаляет запись пива по идентификатору.

#### func  DeleteSnack

```go
func DeleteSnack(db *gorm.DB, id uint) error
```
DeleteSnack удаляет запись закуски по идентификатору.

#### func  UpdateBeer

```go
func UpdateBeer(db *gorm.DB, beer *Beer) error
```
UpdateBeer обновляет существующую запись пива в базе данных.

#### func  UpdateSnack

```go
func UpdateSnack(db *gorm.DB, snack *Snack) error
```
UpdateSnack обновляет существующую запись закуски в базе данных.

#### type Beer

```go
type Beer struct {
	ID        uint           `gorm:"primaryKey"` // Уникальный идентификатор
	CreatedAt time.Time      // Время создания записи
	UpdatedAt time.Time      // Время последнего обновления записи
	DeletedAt gorm.DeletedAt `gorm:"index"` // Время удаления записи (soft delete)

	Name        string  `gorm:"type:varchar(100);not null"` // Название пива
	Brewery     string  `gorm:"type:varchar(100)"`          // Пивоварня
	Style       string  `gorm:"type:varchar(50)"`           // Стиль пива
	Alcohol     float32 `gorm:"type:float"`                 // Содержание алкоголя (%)
	Description string  `gorm:"type:text"`                  // Описание пива
	IBU         int     `gorm:""`                           // Горечь (International Bitterness Units)
	EBC         int     `gorm:""`                           // Цвет (European Brewery Convention)
}
```

Beer представляет модель пива с основными характеристиками.

#### func  GetAllBeers

```go
func GetAllBeers(db *gorm.DB) ([]Beer, error)
```
GetAllBeers возвращает список всех пивных записей.

#### func  GetBeerByID

```go
func GetBeerByID(db *gorm.DB, id uint) (*Beer, error)
```
GetBeerByID возвращает пиво по его идентификатору.

#### type Snack

```go
type Snack struct {
	ID        uint           `gorm:"primaryKey"` // Уникальный идентификатор
	CreatedAt time.Time      // Время создания записи
	UpdatedAt time.Time      // Время последнего обновления записи
	DeletedAt gorm.DeletedAt `gorm:"index"` // Время удаления записи (soft delete)

	Name        string `gorm:"type:varchar(100);not null"` // Название закуски
	Type        string `gorm:"type:varchar(50)"`           // Тип закуски
	Description string `gorm:"type:text"`                  // Описание закуски
	Country     string `gorm:"type:varchar(50)"`           // Страна происхождения
	Calories    int    `gorm:""`                           // Калорийность
	Spicy       bool   `gorm:""`                           // Острая ли закуска
	Vegetarian  bool   `gorm:""`                           // Вегетарианская ли закуска
}
```

Snack представляет модель закуски с основными характеристиками.

#### func  GetAllSnacks

```go
func GetAllSnacks(db *gorm.DB) ([]Snack, error)
```
GetAllSnacks возвращает список всех закусок из базы данных.

#### func  GetSnackByID

```go
func GetSnackByID(db *gorm.DB, id uint) (*Snack, error)
```
GetSnackByID возвращает закуску по её идентификатору.

---

## Пакет: request
# request
--
    import "."

Package request предоставляет удобный обёртку для работы с HTTP-запросами.

Package request предоставляет удобный обёртку для работы с HTTP-запросами.

## Usage

#### type Request

```go
type Request struct {
	Req *http.Request
}
```

Request оборачивает стандартный http.Request и добавляет дополнительные методы
для удобства.

#### func  InitRequest

```go
func InitRequest(r *http.Request) *Request
```
InitRequest создаёт новый экземпляр Request на основе стандартного http.Request.

#### func (*Request) All

```go
func (r *Request) All() url.Values
```
All возвращает все параметры запроса из URL и тела формы в виде url.Values.

#### func (*Request) BearerToken

```go
func (r *Request) BearerToken() string
```
BearerToken извлекает токен Bearer из заголовка Authorization.

#### func (*Request) File

```go
func (r *Request) File(name string) (*multipart.FileHeader, error)
```
File возвращает первый файл с указанным именем из multipart-формы.

#### func (*Request) FullUrl

```go
func (r *Request) FullUrl() string
```
FullUrl возвращает полный URL запроса с протоколом и хостом.

#### func (*Request) HasFile

```go
func (r *Request) HasFile(name string) bool
```
HasFile проверяет, был ли загружен файл с указанным именем.

#### func (*Request) HasHeader

```go
func (r *Request) HasHeader(key string) bool
```
HasHeader проверяет наличие заголовка с указанным ключом.

#### func (*Request) Header

```go
func (r *Request) Header(key string, defaultValue string) string
```
Header возвращает значение заголовка с ключом key, или defaultValue, если
заголовок отсутствует.

#### func (*Request) Json

```go
func (r *Request) Json(v any) error
```
Json декодирует JSON-тело запроса в переданную структуру v.

#### func (*Request) Method

```go
func (r *Request) Method() string
```
Method возвращает HTTP-метод запроса.

#### func (*Request) Path

```go
func (r *Request) Path() string
```
Path возвращает путь запроса.

#### func (*Request) Query

```go
func (r *Request) Query(key string, defaultValue string) string
```
Query возвращает значение параметра query с ключом key, или defaultValue, если
параметр отсутствует.

#### func (*Request) Url

```go
func (r *Request) Url() string
```
Url возвращает URI запроса (путь + query string).

---

## Пакет: router
# router
--
    import "."

Package router реализует простой HTTP-маршрутизатор с поддержкой параметров пути
и методами регистрации обработчиков.

Package router реализует простой HTTP-маршрутизатор с поддержкой параметров
пути.

Package router реализует маршрутизацию HTTP-запросов с поддержкой параметров
пути и JSON-ответов.

## Usage

#### func  GetParams

```go
func GetParams(r *http.Request) map[string]string
```
GetParams извлекает параметры маршрута из контекста HTTP-запроса. Если параметры
отсутствуют, возвращает пустую карту.

#### func  JsonHandlerWrapper

```go
func JsonHandlerWrapper(handler types.JsonHandlerFunc) types.HandlerFunc
```
JsonHandlerWrapper оборачивает JsonHandlerFunc в стандартный http.HandlerFunc,
обеспечивая парсинг параметров, инициализацию запроса и отправку JSON-ответа.

#### func  RegisterRoute

```go
func RegisterRoute(r *Router, path string, handler any, method string)
```
RegisterRoute регистрирует маршрут с указанным HTTP-методом и обработчиком типа
JsonHandlerFunc. Паника происходит, если переданный обработчик не соответствует
типу JsonHandlerFunc.

#### type Route

```go
type Route struct {
}
```

Route описывает один маршрут с шаблоном пути, сегментами и обработчиками по
HTTP-методам.

#### type Router

```go
type Router struct {
}
```

Router содержит список маршрутов и внутренний HTTP-мультиплексор.

#### func  InitRouter

```go
func InitRouter(routes func(r *Router)) (*Router, error)
```
InitRouter инициализирует новый Router, регистрирует маршруты через переданную
функцию routes, и возвращает готовый к использованию маршрутизатор.

#### func (*Router) Connect

```go
func (r *Router) Connect(path string, handler types.JsonHandlerFunc)
```
Connect регистрирует обработчик для HTTP-метода CONNECT по указанному пути.

#### func (*Router) Delete

```go
func (r *Router) Delete(path string, handler types.JsonHandlerFunc)
```
Delete регистрирует обработчик для HTTP-метода DELETE по указанному пути.

#### func (*Router) Get

```go
func (r *Router) Get(path string, handler types.JsonHandlerFunc)
```
Get регистрирует обработчик для HTTP-метода GET по указанному пути.

#### func (*Router) Head

```go
func (r *Router) Head(path string, handler types.JsonHandlerFunc)
```
Head регистрирует обработчик для HTTP-метода HEAD по указанному пути.

#### func (*Router) Options

```go
func (r *Router) Options(path string, handler types.JsonHandlerFunc)
```
Options регистрирует обработчик для HTTP-метода OPTIONS по указанному пути.

#### func (*Router) Patch

```go
func (r *Router) Patch(path string, handler types.JsonHandlerFunc)
```
Patch регистрирует обработчик для HTTP-метода PATCH по указанному пути.

#### func (*Router) Post

```go
func (r *Router) Post(path string, handler types.JsonHandlerFunc)
```
Post регистрирует обработчик для HTTP-метода POST по указанному пути.

#### func (*Router) Put

```go
func (r *Router) Put(path string, handler types.JsonHandlerFunc)
```
Put регистрирует обработчик для HTTP-метода PUT по указанному пути.

#### func (*Router) ServeHTTP

```go
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request)
```
ServeHTTP реализует интерфейс http.Handler и передаёт обработку запросов
внутреннему mux.

---

## Пакет: server
# server
--
    import "."

Package server содержит маршруты инициализации HTTP-сервера.

Package server содержит функции инициализации и запуска HTTP-сервера.

## Usage

#### func  Init

```go
func Init()
```
Init инициализирует роутер, регистрирует маршруты и запускает HTTP-сервер на
порту 8000. В случае ошибки регистрации маршрутов или запуска сервера происходит
логирование и завершение работы.

---

## Пакет: types
# types
--
    import "."


## Usage

#### type HandlerFunc

```go
type HandlerFunc func(w http.ResponseWriter, r *http.Request)
```

HandlerFunc определяет тип функции-обработчика HTTP-запросов.

#### type JsonHandlerFunc

```go
type JsonHandlerFunc func(r *request.Request, params map[string]string) JsonResponse
```

JsonHandlerFunc определяет тип функции-обработчика, которая принимает кастомный
запрос и параметры маршрута, возвращая JSON-ответ.

#### type JsonResponse

```go
type JsonResponse struct {
	Status  string `json:"status"`            // Статус ответа, например "success" или "error"
	Message string `json:"message,omitempty"` // Сообщение об ошибке или дополнительная информация
	Data    any    `json:"data,omitempty"`    // Данные ответа (может быть любого типа)
}
```

JsonResponse представляет структуру JSON-ответа API.

---

