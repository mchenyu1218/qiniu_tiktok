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
            <button class="btn3" >
            </button>
        </div>
        <div class="starbutton">
            <!-- @click="starButton" -->
            <button class="btn4" >
            </button>
        </div>
        <div class="reviewbutton">
            <!-- @click="reviewButton" -->
            <button class="btn5" >
            </button>
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
        methods:{
            upButton(){
                this.$bus.$emit('up')
            },
            downButton(){
                this.$bus.$emit('down')
            }
        }
    }
</script>


<style scoped>
    .buttonlists {
        /* background-color: grey; */
        width: 30px;
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

</style>