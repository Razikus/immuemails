
let formsurl = document.currentScript.getAttribute('formsurl'); 
function vaultFormSubmit (event, form) {
    event.preventDefault()
    let xmlhttp = new XMLHttpRequest(); 
    xmlhttp.onload = () => {
      if (xmlhttp.status >= 200 && xmlhttp.status < 300) {
        const response = JSON.parse(xmlhttp.responseText)
        iziToast.success({
            title: 'OK',
            position: 'topRight',
            message: 'Message sent!',
        });
      } else {
        iziToast.warning({
            title: 'Error',
            position: 'topRight',
            message: 'Cannot send message right now. Please try again later.',
        });
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