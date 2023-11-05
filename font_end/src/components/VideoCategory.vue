<template>
    <div class="videodetails">
        <div class="search">
            
            <div class="category_list" v-for="s in searchData" :key="s.id" @click="findcategory(s)">
                <i :class="s.iconname"></i>
                <div class="item-name">{{ s.menuName }}</div>
            </div>
            
        </div>
    </div>
</template>

<script>
    export default {
        name: 'VideoCategory',
        data() {
            return {
                keywrodssearch:'',
                searchData: [
                    {id: 1,menuName: "首页 ",iconname: "el-icon-s-custom",},
                    {id: 2,menuName: "热门视频",iconname: "el-icon-star-on",},
                    {id: 3,menuName: "体育频道",iconname: "el-icon-star-on",},
                    {id: 4,menuName: "音乐频道",iconname: "el-icon-s-order",},
                    {id: 5,menuName: "舞蹈频道",iconname: "el-icon-s-tools",},
                    {id: 6,menuName: "美食频道",iconname: "el-icon-error",},
                    {id: 7,menuName: "时尚频道",iconname: "el-icon-error",},
                    {id: 8,menuName: "旅游频道",iconname: "el-icon-error",},
                    {id: 9,menuName: "更多",iconname: "el-icon-error",},
                ],
            }
        },
        methods:{
            findcategory(value){
                console.log(JSON.stringify(value))
            },
            keywrods(){
                console.log(this.keywrodssearch)
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .videodetails{
        height: 92vh;
        width:41px;
        float: left;
        /* background-color: rgb(233, 248, 249); */
        background-color: black;
        /* background-color: #fff; */
        border-left: 1px solid white;
    }

    .search{
        padding-top: 10px;
    }

    .h1 {
        color: white;
    }

    /* .el-input {
        margin-top: 10px;
        margin-left: 8px;
    } */

    .category_list {
        /* margin-left: 5px; */
        display: inline-block;
        width: 40px;
        height: 50px;
        margin-top: 5px;
        /* border: 1px solid purple; */
        cursor: pointer;
    }


    .category_list .item-name{
        /* background-color: #fff; */
        /* padding-left: 4px; */
        margin: 0;
        font-size: 8px;
        text-align: center;
        width: 94%;
        /* margin-top: -10px; */
        color: grey;

    }
    .category_list i{
        position: relative;
        color: grey;
        left:5px;
        font-size: 28px;
        /* background-color: skyblue; */
    }


</style>