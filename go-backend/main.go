package main

import ( "encoding/json" "fmt" "net/http" "log" )

type VendorUpdate struct { OrderID string json:"order_id" Status string json:"status" Carrier string json:"carrier" }

func statusWebhookHandler(w http.ResponseWriter, r *http.Request) { if r.Method != http.MethodPost { http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) return }

var update VendorUpdate
err := json.NewDecoder(r.Body).Decode(&update)
if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

// Logic to forward to n8n or update database directly
fmt.Printf("Received update for Order %s: %s\n", update.OrderID, update.Status)
w.WriteHeader(http.StatusOK)


}

func main() { http.HandleFunc("/api/v1/vendor-update", statusWebhookHandler) log.Println("MSP Automation Engine running on :8080") log.Fatal(http.ListenAndServe(":8080", nil)) }