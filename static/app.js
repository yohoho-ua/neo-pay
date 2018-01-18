/**
 * Created by zim on 17-Jan-18.
 */
new Vue ({
        el: '#app',
        // mounted: function () {
        //     this.fetchAssets();
        // },
        data: {
            address: '',
            sellerAddress: '',
            message: 'Click for new payment address',
            balance: -1,
            statusPaidMessage: '',
            checked: false
        },
        methods:{
            checkStatus: function () {
                this.checked = true;
                console.log(this.checked)
                fetch("/status")
                    .then(response => response.json())
                .then(json => {
                    console.log(json);
                this.address = json.address;
                this.statusPaid = json.status;
                this.balance = json.balance;
                if (this.statusPaid){
                    this.statusPaidMessage = "Your payment was made"}
                else {this.statusPaidMessage = "pending..."}
            })
            },

            getAddress: function () {
                fetch("/address")
                    .then(response => response.json())
                .then(json => {
                    console.log(json);
                this.address = json.address;
            })
            }

        }


    }
);