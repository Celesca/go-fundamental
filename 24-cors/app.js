fetch(`https://204c-49-0-91-87.ngrok-free.app/course`)
.then(function (response) {
    return response.json();
})
.then(function (data) {
    appendData(data);
})
.catch(function (err) {
    console.log('error: ' + err);
})

function appendData(data) {
    console.log(data)
    var mainContainer = document.getElementById("myData");
    for (var i = 0; i < data.length; i++) {
        var div = document.createElement("div");
        div.innerHTML = 'ID :' + data[i].ID + ' Name: ' + data[i].Name + ' Price : ' + data[i].Price + ' Instructor : ' + data[i].Instructor;
        mainContainer.appendChild(div);
    }
}