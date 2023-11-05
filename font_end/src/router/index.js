//专门创建整个应用路由器
import VueRouter from 'vue-router'
import CategoryPage from '../components/CategoryPage'
import ShowandCate from '../components/ShowandCate'
// import VideoCategory from '../components/VideoCategory'
import ListSetting from '../components/ListSetting'
import PersonDetails from '../components/PersonDetails'
import PublishVideo from '../components/PublishVideo'
import CategoryPageuser from '../components/users/CategoryPageuser'
import UserFanandlike from '../components/users/UserFanandlike'
import ListSettinguser from '../components/users/ListSettinguser'

export default new VueRouter({
    routes:[
        
        {
            path:'/categorypage',
            component:CategoryPage
        },
        {
            path:'/showandcate',
            component:ShowandCate
        },
        // {
        //     path:'/listsetting',
        //     component:ListSetting,
        //     // children:[
        //     //     {//  分类
        //     //         path:'categorypage',
        //     //         component:CategoryPage
        //     //     },
        //     //     {// 首页
        //     //         path:'showandcate',
        //     //         component:ShowandCate
        //     //     }
        //     // ]
        // },
        {
            path:'/persondetails',
            component:PersonDetails
        },
        {
            path:'/publishvideo',
            component:PublishVideo
        },
        {
            path:'/categorypageuser',
            component:CategoryPageuser
        },
        {
            path:'/userfanandlike',
            component:UserFanandlike
        },
        {
            name:'userpage',
            path:'/listsetting',
            component:{
                default:ListSetting,
                left:ListSettinguser
            }
        },
    ]
})