<template>
    <div class="videodetailsright">
        <div class="blank">
            <div class="publishuser">
                <div class="pubname">
                    <img src="../assets/light.jpg">
                    <span class="pubusername">moyang</span> 
                    <button class="pubstar">关注</button>
                </div>
            </div>  
        </div>
        <div class="reviewsum">评论(1)</div>
        <div class="reviewlists">
            <div class="review">
                <div class="reviewitem">
                    <div class="reviewname"><img src="../assets/light.jpg"><span>02nv</span> <span class="reviewdate">11-1</span></div>
                    <div class="reviewcontent">打的真好真妙！棒棒棒</div>
                </div>
                <hr />
            </div>
            
            
        </div>

    </div>

</template>

<script>
    export default {
        name: 'VideoDetailsright',
    }
</script>



<style scoped>
    .videodetailsright {
        background-color: black;
        height: 92vh;
        width: 22vw;
        float: left;
    }

    .blank {
        height: 50px;
        /* background-color: grey; */
    }

    .publishuser{
        line-height: 50px;
        vertical-align: middle;
        padding-left: 10px;
        
    }

    .publishuser .pubname img{
        width: 39px;
        height: 39px;
        border-radius: 50%;
        margin-right: 10px;
        /* background-color: yellow; */

    }

    .reviewlists{
        height: 87%;
        /* background-color: aquamarine; */
        overflow: auto;
    }

    .reviewsum{
        width: 100%;
        font-size: 12px;
        margin:0 ;
        text-align: center;
        margin-top: 4px;;
        /* background-color: #fff; */
        color: white;
        border-bottom: 1px solid white;
    }

    .pubusername{
        color: white;
    }

    .pubstar{
        margin-left: 10px;
        /* border:0; */
        color: white;
        background-color: black;
        border: 1px solid white;

    }


    .reviewlists::-webkit-scrollbar{
        display: none;
    }


    .review {
        width: 99%;
        margin-top: 10px;
        margin: 0 auto;
        background-color: black;
        
    }

    img{
        width: 20px;
        height:20px;
        border-radius: 50%;
        margin-right: 8px;
        vertical-align: middle;
    }
    .reviewname span{
        vertical-align: middle;
        color: white;
    }

    .reviewname{
        font-size: 12px;
        margin-bottom: 10px;
        
    }

    .reviewdate {
        color: grey;
        font-size: 8px;
    }

    .reviewcontent {
        text-indent: 2em;
        font-size: 10px;
        color: white;
    }

    .reviewitem{
        width: 90%;
        margin: 0 auto;
        margin-top: 5px;
    }
    hr{
        /* color: grey; */
        opacity: 0.2;
    }
</style>