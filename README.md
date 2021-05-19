# Woocommerce rest api client

Questo repository contiene dei metodi che consentono di interfacciarsi a woocommerce attraverso la sua API REST.

## Inizializzazione della libreria.

Importate la libreria con go get all'interno del vostro progetto Go.

Ricordatevi di impostare le seguenti variabili di ambiente:

```
consumerKey := os.Getenv("WC_CONSUMER_KEY")
consumerSecret := os.Getenv("WC_CONSUMER_SECRET")
wpBaseUrl = os.Getenv("WP_BASE_URL")
```

I valori delle prime due sono recuperabili dal pannello di amministrazione di woocommerce (frugate dentro wordpress) ed il terzo è il base url di wordpress.


## Metodi

I metodi attualmente implementati sono:
-  GetProduct(productId) - recupera il relativo oggetto WooCommerceProduct attraverso l'id.