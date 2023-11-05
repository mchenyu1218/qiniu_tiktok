<template>
    <div class="fansorlike">
        <!-- <h1 style="color:white">dadasd</h1> -->
        <div class="fanusers" v-for="item in mydetail" :key="item.id">
                <hr>
            <div class="fanuserimg">
                <img src="../../assets/light.jpg" alt="">
            </div>
            <div class="fanuserdatil">
                <span class="fanname">{{ item.name }}</span><br>
                <span class="fantitle">{{ item.title }}</span>
            </div>
            <div class="fanlike">
                <button>{{item.islike?'已关注':'关注'}}</button>
            </div>
        </div>
    </div>

</template>


<script>
    export default{
        name:'FansandLikedetails',
        props:['mydetail'],
        
    }

</script>

<style>
    .fansorlike .fanusers hr{
        opacity: 0.5;
    }
    .fansorlike{
        width: 84vw;
        /* background-color: grey; */
        height: 80vh;
        overflow: auto;
    }

    .fansorlike::-webkit-scrollbar{
        display: none;
    }

    .fanusers{
        height: 50px;
        /* background-color: skyblue; */
        margin-bottom: 15px;
        margin-top: 15px;
    }

    .fanusers .fanuserimg{
        width: 50px;
        float: left;
        display: inline-block;
    }

    .fanusers .fanuserimg img{
        width: 40px;
        height:40px;
        border-radius: 50%;
    }
    .fanusers .fanuserdatil{
        /* background-color: pink; */
        width: 68vw;
        height: 50px;
        float: left;
        display: inline-block;
        padding-left:15px;
    }

    .fanusers .fanuserdatil .fanname{
        color: white;
        font-size: 20px;
    }

    .fanusers .fanuserdatil .fantitle{
        color: white;
        font-size: 15px;
        opacity: 0.6;
        padding-left: 20px;
        display: inline-block;
        width: 60vw;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }

    .fanusers .fanlike{
        display: inline-block;
        float: left;
    }

    .fanusers .fanlike button{
        margin-top:  10px;
        margin-left: 20px;
        border-radius: 30px;;
        border: 1px solid white;
        color: white;
        background-color: black;
    }
</style>