package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job      // JobQueue (Cola de trabajos)
	WorkerPool chan chan Job // es un canal de canales
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {

	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}

}

func (w Worker) Start() {

	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d \n", w.Id, fib)
			case <-w.QuitChan: // para cuando reciba que tiene que cerrarce el canal.
				fmt.Printf("Worker with id %d Stopped \n", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {

	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {

	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// cremos el metodo para que encole los Jobs
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:

			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job

			}()
		}
	}
}

func (d *Dispatcher) Run() {

	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start() //  les doamos inicio
	}

	go d.Dispatch()

}

// esta funcion es la encargada de manejar todo lo que venga en las peticiones HTTP del Servidor
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {

	if r.Method != "POST" { // GET, PUT, DELETE si es alguno de estos metodos  vamos a manejar el error a continuacion

		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed) // esto le respondera al cliente que no podra utilizar este metodo

	}

	//FormValue nos permite acceder a todo lo que se esta enviando en el Request y de esta forma son capaz de extraerlo
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}
	//con FormValue  nos permite acceder a todo lo que se va a estar enviando
	//en el r *http.Request y validamos si la forma en que se envia la data es la correcta
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Ivalid Name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	//Lo encolamos con el canal de jobQueue
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)

}

func main() {

	const (
		maxWorkers   = 4       // Definimos la cantidad de Worker con los que vamos a trabajar
		maxQueueSize = 20      // Definimos la maxima cantidad de Jobs que van a ser procesados simultaneamente
		port         = ":8081" // definimos por cual puerto nuestro servidor va a estar escuchando
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()
	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))

}
