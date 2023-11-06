<template>
    <div class="pagelists">
        <div class="navigationbar">
            <span @click="changestyle(true)" class="first" :style="styles" >我关注的</span>
            <span  @click="changestyle(false)" class="second" :style="styles2">关注我的</span>
            <!-- <div v-for="cate in catepages" :key="cate.id" class="navbar" >
                <button>{{ cate.menuName }}</button>
            </div> -->
        </div>
        <!-- <hr> -->
        <FansandLikedetails :mydetail="isfanorlike?mylike:myfan" ></FansandLikedetails>
    </div>
</template>


<script>
    import FansandLikedetails from './FansandLikedetails.vue'
    export default{
        name:'UserFanandlike',
        components:{FansandLikedetails},
        data() {
            return {
                styles:{
                    opacity:'1'
                    // border:'1px solid white',
                    // borderBottom:'0',
                    // backgroundColor:'rgba(135, 133, 133, 0.3)',
                    // boxShadow:'0px 0px 10px 10px rgba(135, 133, 133, 0.3)',
                },
                styles2:{
                    opacity:'0.6'
                    // backgroundColor:'black',
                    // boxShadow:'0',
                    // border:'0',
                    // borderBottom:'1px solid black',
                },
                isclick:true,
                mylike:[
                    {id:'1',name:'02nv',title:'这人什么都没说这人什么都没说这人什么都没说这人什么都没说这人什么都没说这人什么都没说',islike:1}
                ],
                myfan:[
                    {id:'1',name:'moyang',title:'这人什么都没说这人什么都没说这人什么都没说这人什么都没说这人什么都没说这人什么都没说',islike:0}
                ],
                isfanorlike:true,
            }
        },
        methods:{
            changestyle(flag){
                if(flag===true){
                    this.isfanorlike=true,
                    this.styles.opacity='1',
                    this.styles2.opacity='0.6'
                    // this.styles.border='1px solid white',
                    // this.styles.borderBottom='1px solid black',
                    // this.styles.backgroundColor='rgba(135, 133, 133, 0.3)',
                    // this.styles.boxShadow='0px 0px 10px 10px rgba(135, 133, 133, 0.3)',
                    // this.styles2.border='0',
                    // this.styles2.borderBottom='1px solid black',
                    // this.styles2.backgroundColor='black',
                    // this.styles2.boxShadow='0px 0px 0px 0px rgba(255, 255, 255, 0)'
                }else{
                    this.isfanorlike=false,
                    this.styles.opacity='0.6',
                    this.styles2.opacity='1'
                    // this.styles2.border='1px solid white',
                    // this.styles2.borderBottom='1px solid black',
                    // this.styles2.backgroundColor='rgba(135, 133, 133, 0.3)',
                    // this.styles2.boxShadow='0px 0px 10px 10px rgba(135, 133, 133, 0.3)',
                    // this.styles.border='0',
                    // this.styles.borderBottom='1px solid black',
                    // this.styles.backgroundColor='black',
                    // this.styles.boxShadow='0px 0px 0px 0px rgba(255, 255, 255, 0)'
                }
                // this.isclick=!this.isclick
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
        margin-left: 20px;
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

    .navigationbar span{
        color: white;
        /* border:1px solid white; */
        margin-right: 10px;
        /* border-bottom: 0; */
        cursor: pointer;
        padding-left: 10px;
        padding-right: 10px;
        /* background-color: black;
        box-shadow: 0; */
    }

    /* .navigationbar #first{
        border:1px solid white;
        border-bottom: 0;
    }

    .navigationbar span:focus{
        
        border:1px solid white;
        background-color: rgba(135, 133, 133, 0.3);
        box-shadow: 0px 0px 10px 10px rgba(135, 133, 133, 0.3);
        border-bottom: 0;
    } */

    /* .navigationbar span.second{
        color: white;

        border-bottom: 0;
        cursor: pointer;
    } */


</style>