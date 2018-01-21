/**
 * Created by zim on 17-Jan-18.
 */
new Vue({
        el: '#app',
        mounted: function () {
            this.repeat();
        },
        data: {
            address: '',
            sellerAddress: '',
            message: 'Click for new payment address',
            deposit: -1,
            statusMessage: '',
            checked: false,
            sum: 0,
            checkSumField: 'form-control',
            testTickr: 0
        },
        methods: {
            checkStatus: function () {
                this.statusMessage = "pending..."
                this.checked = true;
                console.log(this.checked)
                fetch("/status")
                    .then(response => response.json())
                .then(json => {
                    console.log(json);
                 //this.testTickr++;
                this.address = json.address;
                this.statusPaid = json.status;
                this.deposit = json.deposit;
                if (this.statusPaid) {
                    this.statusMessage = "Your payment was made"
                }
                else {
                    this.statusMessage = "Your payment wasn't found"
                }
            })
            },

            getAddress: function () {
                fetch("/address")
                    .then(response => response.json()
                )
                .
                then(json => {
                    console.log(json);
                this.address = json.address;
            })
            },
            repeat: function () {
                //setTimeout(this.checkStatus, 3000)
                setTimeout(this.showTickr())
           },
            showTickr: function () {
                this.testTickr++;
                console.log(this.testTickr)

            }

        }


    }
);