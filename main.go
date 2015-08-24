package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func prepare(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	var in map[string]string
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, fmt.Sprintf(`{"errors":["%v"]}`, err.Error()), 400)
	}
	return in, err
}

func render(w http.ResponseWriter, res map[string]interface{}) {
	w.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, fmt.Sprintf(`{"errors":["%v"]}`, err.Error()), 500)
	}
}

func handleGenerateCPlus(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGenerateRuby(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGenerateNode(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGeneratePython(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGeneratePHP(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGenerateCSharp(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGenerateObjectiveC(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGenerateJava(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func handleGenerateGo(w http.ResponseWriter, r *http.Request) {
	if in, err := prepare(w, r); err == nil {
		out := ""
		render(w, map[string]interface{}{"source": out})
	}
}

func main() {
	http.HandleFunc("generate/c", handleGenerateC)
	http.HandleFunc("generate/cpp", handleGenerateCPlus)
	http.HandleFunc("generate/ruby", handleGenerateRuby)
	http.HandleFunc("generate/node", handleGenerateNode)
	http.HandleFunc("generate/python", handleGeneratePython)
	http.HandleFunc("generate/php", handleGeneratePHP)
	http.HandleFunc("generate/csharp", handleGenerateCSharp)
	http.HandleFunc("generate/objc", handleGenerateObjectiveC)
	http.HandleFunc("generate/java", handleGenerateJava)
	http.HandleFunc("generate/go", handleGenerateGo)

	http.ListenAndServe(":8000", nil)
}
