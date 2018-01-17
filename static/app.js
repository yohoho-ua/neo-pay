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
        },
        methods:{

            getAddress: function () {
                var transactions = [];
                fetch("/neo")
                    .then(response => response.json())
                .then(json => {
                    console.log(json);
                this.address = json.address;
            })
            }

        }


    }
);