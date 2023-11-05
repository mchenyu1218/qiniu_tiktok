<template>
    <div class="pagelists">
        <div class="navigationbar">
            <div v-for="cate in catepages" :key="cate.id" class="navbar" >
                <button>{{ cate.menuName }}</button>
            </div>
        </div>
        <hr>
        <div class="navigationdetails">
            <div v-for="show in showdetailsvedio" :key="show.id" class="showdetails" >
                <video :src="show.path" @click="showcurvedio(show)"></video>
                <div class="showimganddetail">
                    <img src="../assets/light.jpg" alt="">
                    <span>一只鸡的故事一只鸡的故事一只鸡的故事一只鸡的故事</span>
                </div>
                
            </div>
        </div>
    </div>
</template>


<script>

    export default{
        name:'CategoryPage',
        data() {
            return {
                catepages: [
                    {id: 1,menuName: "热门"},
                    {id: 2,menuName: "体育"},
                    {id: 3,menuName: "音乐"},
                    {id: 4,menuName: "舞蹈"},
                    {id: 5,menuName: "美食"},
                    {id: 6,menuName: "时尚"},
                    {id: 7,menuName: "游戏"},
                    {id: 8,menuName: "搞笑"},
                    {id: 9,menuName: "科技"},
                    {id:10,menuName: "汽车"},
                    {id:11,menuName: "纪录片"},
                    {id:12,menuName: "娱乐"},
                    {id:13,menuName: "电视剧"},
                    {id:14,menuName: "知识"},
                    {id:15,menuName: "咨询"},
                    {id:16,menuName: "生活"},
                    {id:17,menuName: "旅行"},
                    {id:18,menuName: "VLOG"},
                    {id:19,menuName: "其他"},
                ],
                showdetailsvedio:[
                    {id: '001',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
                    {id: '002',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                    {id: '003',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
                    {id: '004',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                    {id: '005',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
                    {id: '006',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                    {id: '007',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
                    {id: '008',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                    {id: '009',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
                    {id: '010',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                    {id: '011',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
                    {id: '012',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                    {id: '013',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
                ]
            }
        },
        methods:{
            showcurvedio(id){
                this.$bus.$emit('showcurvedio',id);
            }
        }
    }


</script>


<style scoped>
    .pagelists{
        height: 70vh;
        width: 85vw;
        float: left;
        /* background-color: #fff; */
        margin-left: 70px;
        margin-top: 30px;
    }

    .pagelists hr{
        margin: 0;
        opacity: 0.3;
        margin-bottom: 2px;
    }

    .navigationbar{
        /* background-color: grey; */
        height: 35px;
        display: flex;
        align-items: center; /* 垂直居中 */
    }

    .navbar{
        margin-left: 1px;
    }

    .navbar button{
        color: white;
        background-color: black;
        font-size: 11px;
        height:20px;
        cursor: pointer;
        border:0;
        margin-right:  3px;
        border: 1px solid grey;
        border-radius: 30%;

    }
    .navbar button:hover{
        border: 1px solid white;
        border-radius: 30%;
    }

    .navigationdetails{
        /* background-color: grey; */
        width: 100%;
        height: 100%;
        /* display: table-cell; */
        /* vertical-align: middle; */
        overflow: auto;
        margin-left: 15px;
        margin-top: 5px;
    }

    .navigationdetails::-webkit-scrollbar{
        display: none;
    }


    .showdetails{
        display: inline-block;
        width: 25%;
        height: 37%;
        /* margin-left: 20px; */
        /* border: 1px solid black; */
        /* background-color: #fff; */
        /* display:table-cell; vertical-align:middle; */
    }
    

    .showdetails video{
        /* margin-top: 20px; */
        text-align: center;
        width: 85%;
        cursor: pointer;
    }

    .navigationdetails .showdetails img{
        width: 15px;
        height:15px;
        border-radius: 50%;
    }

    .showimganddetail{
        /* background-color:skyblue; */
    }

    .showimganddetail span{
        font-size: 9px;
        /* height:30px;  */
        margin-left: 4px;
        margin-top: 5px;
        width: 75%;
        display: inline-block;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        color: white;
        /* background-color: black; */
    }
</style>