<template>
    <div class="context">
        <div class="headertitle">
            <div class="title">HMHTOK</div>
        </div>
        <div class="headersearch">
            <el-input v-model="input" placeholder="请输入内容" class="searchinput"></el-input>
        </div>
        <div class="headersign">
           <!-- <span class="login">登录</span>
            <span class="signin">注册</span> -->
            <el-button type="success" class="signin" style="margin-right:30px;" >退出</el-button>
            <el-button type="danger" class="login" style="margin-right:15px;">hyn</el-button>
        </div>
    </div>
    
</template>


<script>
    export default {
        name: 'HeaderCom',
        data() {
            return {
                input: '',
            }
        },
    }
</script>


<!-- <style>
    .menu{
        left: 24px;
        top: 14px;
        width: 199px;
        height: 40px;
        opacity: 1;
        display: flex;
        justify-content: flex-start;
        align-items: center;

    }

    .title{
        left: 64px;
        top: 5.5px;
        width: 88px;
        height: 29px;
        opacity: 1;
        display: flex;

        /** 文本1 */
        font-size: 14px;
        font-weight: 600;
        letter-spacing: 0px;
        line-height: 28.8px;
        color: rgba(0, 0, 0, 1);
        text-align: left;
        vertical-align: top;

    }

    .search{
        left: 847px;
top: 12px;
width: 460px;
height: 44px;
opacity: 1;
display: flex;
justify-content: flex-start;
align-items: center;


    }

    .searcharea{
        left: 0px;
top: 0.5px;
width: 414px;
height: 43px;
opacity: 1;
border-radius: 4px 0px, 0px, 4px;
background: rgba(229, 229, 229, 1);
display: flex;
justify-content: flex-start;
align-items: center;
padding: 12px 12px 12px 12px;

    }

</style> -->

<style scoped>
    .context {
        height: 5vh;
        width: 100vw;
        background-color: black;
        margin: 0;
        padding: 0;
        display: flex;
        position: fixed;

    }

    .headertitle {
        flex: 1;
        height: 100%;
        /* height: 6vh; */
        /* border: 1px solid red; */
        padding-top: 1vh;
    }

    .headertitle .title{
        margin-left: 1vw;
        color: white;
        font-size: 3vh;
        /* border: 1px solid red; */
    }

    .headersearch {
        /* height: 6px; */
        flex: 1;
        /* border: 1px solid red; */

    }

    .searchinput{
        border-radius:3px;
    }

    .headersearch .el-input{
        height: 97%;
        margin-top: 2px;
        /* width: 100%; */
        /* height: 100px; */
        /* border: 1px double black; */
    }
    >>> .el-input__inner{
        height: 4vh;
        width: 30vw;
        top:0;
        background-color: black;
        color: white;

    }

    .headersign {
        flex: 1;
        text-align: center;

    }

    .headersign .el-button{
        float: right;
        margin-top: 4px;
        color: white;
        padding: 7px;
        font-size: 10px;
        text-align: center;

    }


</style>