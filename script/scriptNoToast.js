
let formsurl = document.currentScript.getAttribute('formsurl'); 
function vaultFormSubmitWithCallback (event, form, callback) {
    event.preventDefault()
    let xmlhttp = new XMLHttpRequest(); 
    xmlhttp.onload = () => {
      if (xmlhttp.status >= 200 && xmlhttp.status < 300) {
        const response = JSON.parse(xmlhttp.responseText)
        if(callback != undefined) {
          callback(response)
        }
      } else {
        if(callback != undefined) {
          callback(undefined)
        }
      }
    }
    let theUrl = formsurl + "/" + form.id;
    xmlhttp.open("POST", theUrl);
    xmlhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    let data = new FormData(form);
    let value = Object.fromEntries(data.entries());
    xmlhttp.send(JSON.stringify(value));
    return false
}