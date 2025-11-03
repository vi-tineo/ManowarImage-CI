package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var ingressosDisponiveis = 100
var evento = "Manowar True Metal Concert"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/reserva", reservaHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Evento}}</title>
		<style>
			body {
				background-color: black;
				color: white;
				text-align: center;
				font-family: Arial, sans-serif;
			}
			@keyframes flame {
				0% { background-position: 0% 50%; }
				50% { background-position: 100% 50%; }
				100% { background-position: 0% 50%; }
			}
			h1 {
				font-size: 84px;
				margin-top: 40px;
				background: linear-gradient(270deg, #C0C0C0, #ff4500, #C0C0C0);
				background-size: 600% 600%;
				-webkit-background-clip: text;
				-webkit-text-fill-color: transparent;
				animation: flame 5s ease infinite;
			}
			.player {
				margin-top: 30px;
			}
			.notice {
				color: #ffcc00;
				font-size: 22px;
				margin-top: 10px;
			}
			form {
				margin-top: 50px;
				font-size: 18px; /* redução mais perceptível */
			}
			input[type="checkbox"] {
				transform: scale(1.3); /* menor que antes */
				margin-left: 8px;
			}
			input[type="number"], select {
				font-size: 18px;
				padding: 8px;
				width: 240px;
			}
			input[type="submit"] {
				font-size: 20px;
				padding: 12px 24px;
				margin-top: 30px;
				background: linear-gradient(to right, #ff4500, #ff6347);
				border: none;
				color: white;
				cursor: pointer;
				box-shadow: 0 0 8px #ff4500;
			}
			input[type="submit"]:hover {
				background: linear-gradient(to right, #ff6347, #ff0000);
				box-shadow: 0 0 16px #ff0000;
			}
		</style>
	</head>
	<body>
		<h1>{{.Evento}}</h1>

		<div class="player">
			<iframe width="640" height="360"
				src="https://www.youtube.com/embed/oo5rP_1k4lo?autoplay=1&mute=1&loop=1&playlist=oo5rP_1k4lo"
				frameborder="0" allow="autoplay; encrypted-media" allowfullscreen>
			</iframe>
			<div class="notice">Click the speaker icon to unleash the sound of True Metal 🤘</div>
		</div>

		<form action="/reserva" method="POST">
			<p>Are you True Metal? <input type="checkbox" name="truemetal"></p>
			<p>Are you a student? <input type="checkbox" name="student"></p>
			<p>How many tickets? <br><input type="number" name="ingressos" min="1" max="{{.Disponiveis}}"></p>
			<p>Payment method:<br>
				<select name="pagamento">
					<option value="cash">Cash</option>
					<option value="credit">Credit Card</option>
				</select>
			</p>
			<input type="submit" value="Book Now">
		</form>
	</body>
	</html>
	`
	t := template.Must(template.New("home").Parse(tmpl))
	t.Execute(w, struct {
		Evento      string
		Disponiveis int
	}{
		Evento:      evento,
		Disponiveis: ingressosDisponiveis,
	})
}

func reservaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trueMetal := r.FormValue("truemetal") == "on"
	student := r.FormValue("student") == "on"
	ingressosStr := r.FormValue("ingressos")
	pagamento := r.FormValue("pagamento")

	ingressos, err := strconv.Atoi(ingressosStr)
	if err != nil || ingressos < 1 || ingressos > ingressosDisponiveis {
		http.Error(w, "Número de ingressos inválido", http.StatusBadRequest)
		return
	}

	if !trueMetal {
		fmt.Fprintln(w, `
			<html><body style="background-color:black; text-align:center; font-family:Arial">
			<h1 style="color:red; font-size:72px; margin-top:120px;">You are not True Metal! Get out!</h1>
			</body></html>
		`)
		return
	}

	ingressosDisponiveis -= ingressos

	preco := 1000.0
	if student {
		preco = 500.0
	}

	fmt.Fprintf(w, `
		<html><body style="background-color:black; color:white; text-align:center; font-family:Arial">
		<h2 style="font-size:48px;">Booking Confirmed!</h2>
		<p style="font-size:28px;">Tickets: %d</p>
		<p style="font-size:28px;">Student: %v</p>
		<p style="font-size:28px;">Payment Method: %s</p>
		<p style="font-size:28px;">Total: $%.2f</p>
		<p style="font-size:32px;">Hail and Kill! 🤘</p>
		</body></html>
	`, ingressos, student, pagamento, preco*float64(ingressos))
}
