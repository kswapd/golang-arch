Vue.component('menu-component', {
    props: {
        items: {
            type: Array,
            required: true
        }
    },
    template: `
    <nav class="menu">
      <ul>
        <li v-for="item in items" :key="item.text">
          <a :href="item.link" @click.prevent="$emit('menu-click', item)">{{ item.text }}</a>
        </li>
      </ul>
    </nav>
  `
});

var app = new Vue({
    el: '.main-container',
    data: {
        message: 'Hello Vue!',
        menuItems: [
            { text: 'Home', link: '/' },
            { text: 'About', link: '/about' },
            { text: 'Contact', link: '/contact' }
        ]
    },
    mounted: function () {
        console.log('dddd')
        this.getData();
    },
    methods: {
        formatDate: function (date, format) {
            const map = {
                'mm': ('0' + (date.getMonth() + 1)).slice(-2),
                'dd': ('0' + date.getDate()).slice(-2),
                'yyyy': date.getFullYear(),
                'HH': ('0' + date.getHours()).slice(-2),
                'MM': ('0' + date.getMinutes()).slice(-2),
                'SS': ('0' + date.getSeconds()).slice(-2),
            };
            return format.replace(/mm|dd|yyyy|HH|MM|SS/gi, matched => map[matched]);
        },
        getData: async function () {
            this.message = "Loading...";
            const response = await fetch('/api/data');
            const data = await response.json();
            console.log(data)
            if (data["messageType"] == "1000") {
                const milliseconds = data["body"]["tradeTime"];
                const date = new Date(milliseconds);
                const formattedDate = this.formatDate(date, 'yyyy-mm-dd HH:MM:SS');
                this.message = "名称: " + data["body"]["name"] + "<br>";
                this.message += "交易时间: " + formattedDate + "<br>";
                this.message += "价格:" + data["body"]["price"] + "(" + data["body"]["unit"] + ")";
            }
        },
        handleMenuClick: function (item) {
            if (item.text === 'About') {
                this.message = '<b>About Page</b><br>This is a personal note and demo site. Here you can find information about the author, project goals, and more.';
            } else if (item.text === 'Home') {
                this.getData();
            } else if (item.text === 'Contact') {
                this.message = '<b>Contact Page</b><br>For inquiries, please email: example@example.com';
            }
        }
    }
});



