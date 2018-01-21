/**
 * Created by zim on 19-Jan-18.
 */

window.App = {
    start: function () {
        window.App.checkStatus();
        window.App.startTrack();
    },
    checkStatus: function () {
        var url = "/status";
        fetch(url)
            .then(response => response.json())
            .then(customer => {
                console.log(customer);
                document.getElementById("address").innerHTML = customer.address;
                document.getElementById("balance").innerHTML = customer.deposit;
                document.getElementById("checked").style.display = "block";
                document.getElementById("status").innerHTML = "pending...";
                document.getElementById("blockcount").innerHTML = (customer.block)-1;

            })
    },
    startTrack : function() {
        setInterval(function () {
            // console.log("!!!")
            window.App.checkStatus()
        }, 10000);
    },

    getNewAddress:function () {
        $.get("/status?new=true");
    }
}

window.addEventListener('load', function () {
    App.start()
})