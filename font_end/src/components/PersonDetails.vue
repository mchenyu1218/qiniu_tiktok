<template>
    <div class="perdetails">
        <div class="personimg">
            
            <div class="userimg">
                <img src="../assets/searchicon/dance.png" alt="">
            </div>
            <div class="personleft">
                <div class="usernname">{{userinfo?.name}}</div>
                <div class="userdetail">这个人什么都没有说这个人什么都没有说这个人什么都没有说这个人什么都没有说这个人什么都没有说</div>
            </div>
            <div class="personright">
                <span class="userpro">作品数：0</span>
                <span class="userstar">关注数：{{ userinfo?.follow_count  }}</span>
                <span class="userfans">粉丝数：{{userinfo?.follower_count}}</span>
                <span class="userlike">收藏数：0</span>
                <button class="addnew" @click="pubnewvideo">点击创作<i class="el-icon-edit"></i> </button>
            </div>
        </div>
        <!-- 最近浏览 -->
        <div class="personrecently ">
            <span class="span1">最近浏览-></span>
            <span class="span2">更多>></span>
            <div class="undercontext">
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/b868152d39f68221162ed58d112bbf76.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/7621a726f3df74c6b9bec19aefdc7407.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4" ></video>
                </div>
            </div>
        </div>
        <!-- 收藏 -->
        <div class="personstar">
            <span class="span1">收藏-></span>
            <span class="span2">更多>></span>
            <div class="undercontext">
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/b34f19fc522fc7ef9f86d28ffa7cb0d7.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/7621a726f3df74c6b9bec19aefdc7407.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/b868152d39f68221162ed58d112bbf76.mp4" ></video>
                </div>
            </div>
        </div>
        <!-- 点赞 -->
        <div class="personlike">
            <span class="span1">点赞-></span>
            <span class="span2">更多>></span>
            <div class="undercontext">
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/7621a726f3df74c6b9bec19aefdc7407.mp4" ></video>
                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4" ></video>

                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/b34f19fc522fc7ef9f86d28ffa7cb0d7.mp4" ></video>

                </div>
                <div class="contextitem">
                    <video src="http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4" ></video>
                </div>
            </div>
        </div>


    </div>
</template>

<script>
    import {getuser} from '../api/vedio'
    export default {
        name: 'PersonDetails',
        data() {
            return {
                userid:'1672489513669',
                // follow_count:0,
                // follower_count:0,
                // name:'',
                userinfo:{}
            }
        },
        computed: {

        },
        methods: {
            pubnewvideo(){
                console.log('sadfa');
                this.$bus.$emit('newvediopage',false);
                // this.$bus.$emit('showpage',param);
            }
        },
        async mounted(){
            const res=await getuser(this.userid,this.userid)
            const userInfo=res.data.User
            console.log('userinfo=',userInfo);
            console.log('userinfo..=',userInfo.follower_count);
            this.userinfo=userInfo
            // this.userinfo=[...this.userinfo,...userInfo]
            console.log(this.userinfo);
            // this.showdetailsvedio=tagres.data.video_list

            // await this.$axios({
            //     method: "get",
            //     url: "http://localhost:8080/api/tiktok/user/search/:UserID"+newVal,
            //     })
            //     .then((reponse) =>{
            //         console.log("ta=",this.target);
            //         const res=JSON.parse(JSON.stringify(reponse.data.video_list))
            //         this.showdetailsvedio=res
            //         console.log(res);
            //     })
            //     .catch((error)=>{
            //         console.log(error.message);
            // })
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .perdetails {

        height: 85vh;
        width: 75vw;
        background-color: rgba(135, 133, 133, 0.3);
        /* border: 2px solid grey; */
        display: flex;
        flex-direction: column;
        margin-left: 120px;
        /* margin-top: 20px; */
        border-radius: 20px;
        box-shadow: 0px 0px 50px 50px rgba(135, 133, 133, 0.3);

    }

    .personimg {
        flex: 1;
        /* border: 1px solid black; */
        /* border-radius: 50%; */
        
    }
    .userimg{
        /* text-align: center; */
        height: 90%;
        float: left;
        /* margin-bottom: 20px; */
        /* border: 0; */

    }


    /* video{
        width: 100%;
        height: 100%;
    } */

    img {
        height:110px;
        width: 110px;
        margin-top: 10px;
        border-radius: 50%;
        border: 4px  double  gold;
    }

    .personleft{
        float: left;
        margin-left: 30px;
        margin-top: 20px;
        width: 40%;
        /* background-color: yellow; */
    }

    .personleft .usernname{
        color: white;
        font-weight: bold;
        font-size: 20px;
    }

    .personleft .userdetail{
        color: white;
        font-size: 15px ;
        opacity: 0.5;
        /* padding-left: 50px; */
        text-indent: 2em;
    }

    .personright{
        float: right;
        width: 40%;
        padding-top: 25px;
        font-size: 15px;
    }

    .personright span{
        padding-right: 10px;
        color: white;
        opacity: 0.8;
    }

    .personright .addnew{
        cursor: pointer;
        display: block;
        margin-top: 20px;
        margin-left:70px;
        font-size: 25px;
        border-radius: 20px;
        padding-left: 15px;
        padding-right: 15px;
    }

    .personrecently {
        flex: 1;
        /* border-top: 2px solid rgb(200, 228, 255); */
    }

    .personstar {
        flex: 1;
        /* border-top: 2px solid rgb(200, 228, 255); */

    }

    .personlike {
        flex: 1;
        /* border-top: 2px solid rgb(200, 228, 255); */

    }

    .span1 {
        font-weight: bold;
        margin-left: 10px;
        color: white;
    }

    .span2 {
        font-weight: bold;
        margin-right: 10px;
        color: white;
        float: right;
    }

    .span2:hover {
        text-decoration: underline;
    }

    .undercontext {
        height: 80%;
        width: 98%;
        margin: 0 auto;
        /* background-color: rgb(244, 246, 246); */
        display: flex;
        justify-content: space-between;
    }

    .undercontext .contextitem video{
        width: 100%;
        height: 100%;
        margin: 0;
    }

    .contextitem {
        /* border:1px solid */
        display: inline-block;
        margin: auto;
        width: 100%;
        height:100px;
        /* background-color: rgb(200, 205, 205); */

    }
</style>