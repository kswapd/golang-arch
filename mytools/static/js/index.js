


function addScript(url) {
    document.write("<script language=javascript src=" + url + "></script>");
}
const About = { template: `<div><h1>About Us</h1><p>Learn more about our team and company.</p></div>` }
const Contact = { template: `<div><h1>Contact Us</h1><p>Get in touch with us.</p></div>` }
const Team = { template: `<div><h1>Our Team</h1><p>Meet our amazing team.</p></div>` }
const Company = { template: `<div><h1>Our Company</h1><p>Learn about our company's history and mission.</p></div>` }
const WebDevelopment = { template: `<div><h1>Web Development</h1><p>Details about our web development services.</p></div>` }
const MobileDevelopment = { template: `<div><h1>Mobile Development</h1><p>Details about our mobile development services.</p></div>` }
const Services = { template: `<div><h1>Our Services</h1><p>Explore the services we offer.</p></div>` }

const routes = [
    { path: '/', component: Home },
    { path: '/home', component: Home },
    { path: '/about/team', component: Team },
    { path: '/about/company', component: Company },
    { path: '/services/web-development', component: WebDevelopment },
    { path: '/services/mobile-development', component: MobileDevelopment },
    { path: '/contact', component: Contact }
]

// 3. 创建 router 实例，然后传 `routes` 配置
// 你还可以传别的配置参数, 不过先这么简单着吧。
const router = new VueRouter({
    routes // (缩写) 相当于 routes: routes
})


var app = new Vue({
    router,
    el: '.main-container',
    data: {
        splitChar: '#',
        pathChar: '/',
        message: 'Hello Vue!',
        menuItems: [
            {
                text: 'Home',
                url: '/',
                isOpen: false,
                children: []
            },
            {
                text: 'About',
                url: '/about', // URL for About
                isOpen: false,
                children: [
                    { text: 'Team', url: '/about/team', isOpen: false, children: [] },
                    { text: 'Company', url: '/about/company', isOpen: false, children: [] }
                ]
            },
            {
                text: 'Services',
                url: '/services', // URL for Services
                isOpen: false,
                children: [
                    { text: 'Web Development', url: '/services/web-development', isOpen: false, children: [] },
                    { text: 'Mobile Development', url: '/services/mobile-development', isOpen: false, children: [] }
                ]
            },
            {
                text: 'Contact',
                url: '/contact', // URL for Contact
                isOpen: false,
                children: []
            }
        ],
    },
    methods: {
        switchView: function (view) {
            // window.location.hash = this.splitChar + view.toLowerCase();
            console.log("window.location.hash:", window.location.hash);
            console.log("window.location.pathname:", window.location.pathname);
        },
        toggleMenu: function (item) {
            item.isOpen = !item.isOpen;
        },

        handleMenuClick(item) {
            if (item.children && item.children.length > 0) {
                // 如果有子菜单，只展开/折叠，不跳转
                item.isOpen = !item.isOpen;
            } else {
                // 如果没有子菜单，跳转页面
                if (this.$route.path !== item.url) {
                    this.$router.push(item.url);
                }
            }
        },
        // 自动展开当前路径对应的父菜单
        updateMenuOpenState() {
            this.menuItems.forEach(item => {
                if (item.children && item.children.some(child => this.$route.path === child.url)) {
                    item.isOpen = true;
                } else {
                    item.isOpen = false;
                }
            });
        }
    },
    watch: {
        '$route': {
            immediate: true,
            handler() {
                this.updateMenuOpenState();
            }
        }
    }
});



