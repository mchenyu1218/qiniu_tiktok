<template>
    <div class="buttonlists">
        <div class="upbutton">
            <button class="btn1" @click="upButton">
            </button>
        </div>
        <div class="upbutton">
            <button class="btn2" @click="downButton">
            </button>
        </div>
        <div class="heartbutton">
            <!-- @click="heartButton" -->
            <button class="btn3" @click="heartbutton">
            </button>
            <span class="undernumber">{{ favorite_count }}</span>
        </div>
        <div class="starbutton">
            <!-- @click="starButton" -->
            <button class="btn4" >
            </button>
            <span class="undernumber">{{ collect_count }}</span>
        </div>
        <div class="reviewbutton">
            <!-- @click="reviewButton" -->
            <button class="btn5" >
            </button>
            <span class="undernumber">{{ comment_count }}</span>
        </div>
        <div class="sharebutton">
            <!-- @click="shareButton" -->
            <button class="btn6" >
            </button>
        </div>

    </div>

</template>


<script>
    export default {

        name: 'IconList',
        data() {
            return {
                video_detail:[],
                heart:0,
                like:0,
                share:0
            }
        },
        props:['favorite_count','collect_count','comment_count'],
        methods:{
            upButton(){
                this.$bus.$emit('up')
            },
            downButton(){
                this.$bus.$emit('down')
            },
            heartbutton(){
                
            }
        },
        async mounted(){
            await this.$axios({
            method: "get",
            url: "http://localhost:8080/api/tiktok/feedallvideo",
            })
            .then((reponse) =>{
                const res=JSON.parse(JSON.stringify(reponse.data.video_list))
                this.video_detail=res
                // console.log(this.video_detail);
            })
            .catch((error)=>{
                console.log(error.message);
            })
        }
    }
</script>


<style scoped>
    .buttonlists {
        /* background-color: grey; */
        width: 30px;
    }

    .buttonlists div{
        /* border:1px solid red; */
    }

    .buttonlists .btn1{
        margin-top: 10px;
        width: 30px;
        height: 30px;
        background-image: url('../assets/up-circle.png'); 
        background-size:30px 30px;
        background-color: transparent;
        border: 0;
    }

    .buttonlists .btn2{
        margin-top: 10px;
        width: 30px;
        height: 30px;
        background-image: url('../assets/down-circle.png'); 
        background-size:30px 30px;
        background-color: transparent;
        border: 0;
    }

    .buttonlists .btn3{
        margin-top: 10px;
        width: 30px;
        height: 30px;
        background-image: url('../assets/heart.png'); 
        background-size:30px 30px;
        background-color: transparent;
        border: 0;
    }

    .buttonlists .btn4{
        margin-top: 10px;
        width: 30px;
        height: 30px;
        background-image: url('../assets/star.png'); 
        background-size:30px 30px;
        background-color: transparent;
        border: 0;
    }

    .buttonlists .btn5{
        margin-top: 10px;
        width: 30px;
        height: 30px;
        background-image: url('../assets/review.png'); 
        background-size:30px 30px;
        background-color: transparent;
        border: 0;
    }
    .buttonlists .btn6{
        margin-top: 10px;
        width: 30px;
        height: 30px;
        background-image: url('../assets/share.png'); 
        background-size:30px 30px;
        background-color: transparent;
        border: 0;
    }

    /* .buttonli{
        margin-top: 20px;
    } */
    button:hover{
        cursor:  pointer;
    }


    .undernumber{
        color:white;
        display: inline-block;
        text-align: center;
        width:30px;
        font-size: 10px;
        /* background-color: black; */
        float:left;
    }

</style>