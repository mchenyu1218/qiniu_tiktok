<template>
  <div class="videolist"  >

    <!-- autoplayref="focuson" -->
    <div class="videosrc" >
      <video autoplay loop :src="currentVideoPath" @mouseover="handleMouseOver" @mouseout="handleMouseOut" :controls="ifcontrols" ref="videoElement" @keydown="handleKeyDown"></video>
      <div class="iconlists">
        <IconList />
      </div>
    </div>


    <!-- 视频标题、简介等 -->
    <div class="videodetails">
      <h2>一只鸡</h2>
      <span class="videodetail">这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事这是一只鸡的故事</span>
    </div>

  </div>
</template>

<script>
  import IconList from './IconList.vue'
  export default {
    name: 'VideoList',
    data() {
      return {
        video_src: [
        {id: '001',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
        {id: '002',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
        {id: '003',path: 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4',title: ''},
        {id: '004',path: 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4',title: ''},
        ],
        currentIndex: 0,
        ifcontrols:false,

      }
    },
    computed: {
      currentVideoPath() {
        // console.log(JSON.stringify(this.video_src[this.currentIndex]))
        return this.video_src[this.currentIndex].path
      }
    },
    methods: {
      handleKeyDown(event) {
        console.log(event)
        if (event.key === 'ArrowUp') {
          // 按下了上箭头键
          // 执行相应的逻辑
          if (this.currentIndex > 0) {
            this.currentIndex--;
          }
          // console.log('按下了上箭头键')
        } else if (event.key === 'ArrowDown') {
          // 按下了下箭头键
          // 执行相应的逻辑
          if (this.currentIndex < this.video_src.length - 1) {
            this.currentIndex++;
          }
          // console.log('按下了下箭头键')
        }
      },
      handleMouseOver(){
        this.ifcontrols=true
      },
      handleMouseOut(){
        this.ifcontrols=false
      }
    },
    components: {
      IconList,
    },
    mounted(){
      this.$bus.$on('up',()=>{
        if (this.currentIndex > 0) {
            this.currentIndex--;
          }
      }),
      this.$bus.$on('down',()=>{
        if (this.currentIndex < this.video_src.length - 1) {
            this.currentIndex++;
          }
      }),
      this.$bus.$on('showcurvedio',(showvedio)=>{
        
        // console.log(showvedio);
        const foundIndex=this.video_src.findIndex(item=>item.id===showvedio.id)
        if(foundIndex!==-1) {
          const foundItem = this.video_src.splice(foundIndex, 1)[0];
          this.video_src.unshift(foundItem);
        } else {
          const newData={id:showvedio.id,path:showvedio.path,title:''};
          this.video_src.unshift(newData);
        }
        this.currentIndex=0;
        // console.log(this.video_src);
        this.$bus.$emit('showpage',1);
      })
      
      window.addEventListener('keydown',this.handleKeyDown);
    },
    beforeDestroy(){
      window.removeEventListener('keydown',this.handleKeyDown);
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .videolist {
    height: 92vh;
    width: 100%;
    background-color: black;
  }

  .search {
    /* flex: 1; */
  }

  .videosrc {
    /* background-color: grey; */
    position: relative;
    /* height:88%; */
  }


  .iconlists {
    /* display: inline-block; */
    /* float: left; */
    position: absolute;
    right: 10px;
    top:  15%;
  }

  .videosrc span.before {
    color: black;
    float: left;
    background-color: blue;
  }

  .videosrc span.next {
    color: black;
    float: right;
    background-color: blue;
  }

  video {
    width: 100%;
    border: 0;
    /* margin-top: 50px; */
  }


  .videodetails {
    /* padding-left: 10px; */
    color: white;
    background-color: black;
    /* margin-bottom: 100px; */
    /* flex: 3; */
  }

  .videodetails h2 {
    margin: 0;
    margin-left: 5px;
    margin-bottom: 5px;
  }

  .videodetail {
    display: inline-block;
    margin-left: 50px;
    width: 93%;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    font-size: 12px;
  }
</style>