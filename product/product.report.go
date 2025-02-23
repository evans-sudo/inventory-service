package product

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"path"
	"text/template"
	"time"
)

type ProductReportFilter struct {
	NameFilter string `json:"productName"`
	ManufacturerFilter string `json:"manufacturer"`
	SKUFilter string `json:"sku"`
}



func handleProductReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var productFilter ProductReportFilter
		err := json.NewDecoder(r.Body).Decode(&productFilter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return 
		}

		products, err := searchProductData(productFilter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := template.New("report.gotmpl").Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
		t, err = t.ParseFiles(path.Join("templates", "report.gotmpl"))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}

		w.Header().Set("Content-Disposition", "Attachment")

		var tpl bytes.Buffer
		err = t.Execute(&tpl, products)
		rdr := bytes.NewReader(tpl.Bytes())
		http.ServeContent(w, r, "report.html", time.Now(), rdr)

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}