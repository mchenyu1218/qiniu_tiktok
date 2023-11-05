<template>

    <div class="publishvideo">
        <div class="pushdetails">
            <h2>上传视频</h2>
            <div class="publishform">
                <form action="">
                    <div>
                        <label for="videoname">
                            <span>标题：</span>
                            <input type="text" id="videoname">
                        </label>
                    </div>
                    <div>
                        <label for="videomessage">
                            <span>视频内容简介：</span>
                            <textarea name="videomessage" id="videomessage" cols="30" rows="10"></textarea>
                            <!-- <input type="text" id="videomessage"> -->
                        </label>
                    </div>
                    <div>
                        <label for="videomessage">
                            <span>选择视频：</span>
                            <input type="file" id="videoname">
                        </label>
                    </div>
                    <div>
                        <input type="submit" value="发布" id="submit">  
                    </div>



                </form>
            </div>

        </div>

    </div>
</template>



<script>
    export default {
        name: 'PublishVideo',
    }
</script>



<style>
    .publishvideo {
        height: 85vh;
        width: 75vw;
        background-color: rgba(135, 133, 133, 0.3);
        margin-left: 120px;
        /* margin-top: 20px; */
        border-radius: 20px;
        box-shadow: 0px 0px 50px 50px rgba(135, 133, 133, 0.3);
    }

    .pushdetails h2 {
        color: rgb(255, 255, 255);
        opacity: 0.7;
        margin-left: 30vw;
        margin-top: 0;
    }

    .publishform {
        /* border: 1px solid red; */
        /* margin-top: 20px; */
        margin-left: 20vw;
    }

    .publishform span{
        display: inline-block;
        /* background-color: grey; */
        width:10vw;
        text-align: end;
        color: white;
    }

    .publishform div{
        margin-top: 20px;
    }

    #videomessage{
        vertical-align:top
    }

    #videoname{
        color: white;
    }

    .publishform #submit{
        margin-left: 10vw;
    }
</style>