/**
 * Created by zim on 19-Jan-18.
 */

window.App = {
    start: function () {

    },
    getAddress: function () {
        var url = "/status";
        text = $('#address').text();
        if (text === "") {
            url = "/address"
        }
        fetch(url)
            .then(response => response.json())
            .then(customer => {
                console.log(customer);
                document.getElementById("address").innerHTML = customer.address;
                document.getElementById("balance").innerHTML = customer.deposit;
                document.getElementById("checked").style.display = "block";

                setInterval(function () {
                   // console.log("!!!")
                    window.App.getAddress()
                }, 3000);
            })
    },

    setStatus: function () {
        $.get("/status", function (data, status) {
            $("#balance").text(data.block)
            console.log("Data: " + data.json() + "\nStatus: " + status);
        });
    }
}

window.addEventListener('load', function () {
    App.start()
})