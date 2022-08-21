package adapter

// Target предоставляет интерфейс, который нам нужен
type Target interface {
	Request() string
}

// Adaptee реализует систему для которой мы хотим реализовать интерфейс Target
type Adaptee struct {
}

// NewAdapter Конструктор для создания нового адаптера, который возвращает Target
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest реализация функции
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter  реализует имплементирует интерфейс Target и содержит в себе структуру Adaptee
type Adapter struct {
	*Adaptee
}

/* Request это адаптивный метод, который позволяет использовать для Adaptee интерфейс Target,
с помощью нашего адаптера
*/

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
