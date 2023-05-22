# Immutable forms

## Starting container

Check docker-compose and adjust environment variables

If you need SSL - you either need to add traefik as SSL proxy or provide SSL by your own 

Adding SSL by your own:

```
version: '3.3'
services:
  immuemails:
    ports:
        - '443:443'
    environment:
        - LEDGER=yourledger
        - COLLECTION_NAME=yourcollection
        - API_KEY=yourapikey
        - KEY_FILE=/key
        - CERT_FILE=/cert
    restart: always
    image: razikus/experiments:immuemails
    volumes:
        - "./key:/key"
        - "./cert:/cert"
```

## Converting a form into immutable form (without toasts)

Check example/indexNoToast.html

```
Add script src into your page and point formsurl into your URL with ending /form


<script formsurl="http://localhost:8081/form" src="https://razikus.github.io/immuemails/script/scriptNoToast.js"></script>


Add id= tag and onsubmit= tag into form (you can change last parameter to undefined or your own callback)


<form id="form1" onsubmit="vaultFormSubmitWithCallback(event, this, (resp) => {alert(resp.message)})" method="post">
    <label for="fname">First name:</label>
    <input type="text" id="fname" name="fname"><br><br>
    <label for="lname">Last name:</label>
    <input type="text" id="lname" name="lname"><br><br>
    <input type="submit" value="Submit">
</form>

All parameters would be submitted to the vault

```


## Converting a form into immutable form (with toasts)


Check example/index.html

```
If you want to have toasts integrated (with iziToast):


<script src="https://cdnjs.cloudflare.com/ajax/libs/izitoast/1.4.0/js/iziToast.min.js"
    integrity="sha512-Zq9o+E00xhhR/7vJ49mxFNJ0KQw1E1TMWkPTxrWcnpfEFDEXgUiwJHIKit93EW/XxE31HSI5GEOW06G6BF1AtA=="
    crossorigin="anonymous" referrerpolicy="no-referrer"></script>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/izitoast/1.4.0/css/iziToast.min.css"
    integrity="sha512-O03ntXoVqaGUTAeAmvQ2YSzkCvclZEcPQu1eqloPaHfJ5RuNGiS4l+3duaidD801P50J28EHyonCV06CUlTSag=="
    crossorigin="anonymous" referrerpolicy="no-referrer" />
<script formsurl="http://localhost:8081/form" src="https://razikus.github.io/immuemails/script/script.js"></script>

Add id and onsubmit into form


<form id="form1" onsubmit="vaultFormSubmit(event, this)" method="post">
    <label for="fname">First name:</label>
    <input type="text" id="fname" name="fname"><br><br>
    <label for="lname">Last name:</label>
    <input type="text" id="lname" name="lname"><br><br>
    <input type="submit" value="Submit">
</form>

```