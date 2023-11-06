<template>
  <div class="container">
    <div class="uplist">
      <HeaderCom />
    </div>
    <div class="underlist">
      <div class="leftlist">
        <ListSettinguser v-if="!isuserpage"/>
        <ListSetting   v-if="isuserpage"/>
      </div>

      <div class="home"  v-if="!isRouterAlive" >
        <ShowandCate/>
      </div>
      <router-view v-if="isRouterAlive"/>

      <!-- style="display:none" -->
      <!-- <PublishVideo ></PublishVideo> -->

      <!-- <div style="display: none;">
        <CategoryPage/>
      </div> -->
    </div>

    <!-- <div class="header">
      <HeaderCom />
    </div> -->




  </div>
</template>

<script>
  // import VideoList from './components/VideoList.vue'
  import ListSetting from './components/ListSetting.vue'
  // import VideoCategory from './components/VideoCategory.vue'
  // import VideoDetailsright from './components/VideoDetailsright.vue'
  // import PersonDetails from './components/PersonDetails.vue'
  import HeaderCom from './components/HeaderCom.vue'
  // import CategoryPage from './components/CategoryPage.vue'
  import ShowandCate from './components/ShowandCate.vue'
  // import PublishVideo from './components/PublishVideo.vue'
  import ListSettinguser from './components/users/ListSettinguser.vue'
  export default {
    name: 'App',
    components: {
      HeaderCom,
      // VideoList,
      ListSetting,
      // VideoCategory,
      // VideoDetailsright,
      // CategoryPage,
      // PersonDetails,
      ShowandCate,
      ListSettinguser
      // PublishVideo,
    },
    data() {
      return {
        isRouterAlive:false,
        isuserpage:true
      }
    },
    mounted() {
      this.$bus.$on('showpage', (data) => {
        this.isRouterAlive=true
        // console.log(data)
        if (data == 1||data==6) {
          this.$router.push({
            path: '/showandcate',
          }).catch(err => err)
        } else if (data == 2) {
          this.$router.push({
            path: '/persondetails'
          }).catch(err => err)
        }else if(data==3){
          this.$router.push({
            path: '/categorypage'
          }).catch(err => err)
        }else if(data==4){
          console.log(data);
        }else if(data==='new'){
          this.$router.push({
            path: '/publishvideo'
          }).catch(err => err)
        }else if(data==7){
          this.$router.push({
            path: '/categorypageuser'
          }).catch(err => err)
        }else if(data==8){
          this.$router.push({
            path: '/userfanandlike'
          }).catch(err => err)
        }
      }),
      this.$bus.$on('userpage',(data)=>{
        this.isuserpage=data
        this.$bus.$emit('showpage',6);
      }),
      this.$bus.$on('newvediopage',(data)=>{
        this.isuserpage=data
        this.$bus.$emit('showpage','new');
      })
    },
    
    beforeDestroy(){
        this.$bus.$off('showpage'),
        this.$bus.$off('userpage'),
        this.$bus.$off('newvediopage')
    }
  }
</script>





<style>
  body {
    /* background-image: url('./assets/light.jpg'); */
    background-size: 100%;
    margin: 0px;
    background-color: black;
  }

  .container {
    width: 1890px;
    min-width: 5700px;

    background-color: black;
    position: relative;
  }

  .uplist {
    height: 5vh;
    /* background-color: #fff; */
  }

  .leftlist {
    /* width: 4vw; */
    /* height: 3vh; */
    float: left;
    /* background-color: #fff; */
  }

  /* .middlelist1 {
    height: 92vh;
    width: 75vw;
    float: left;
  }

  .middlelist2 {
  }

  .rightlist {
    height: 92vh;
    width: 5vw;
    float: left;
  } */

  /* .header {
    margin: 0;
    padding: 0;

    right: 0;
    left: 0;
    margin: 0 auto;
    position: absolute;
  } */
</style>


<!-- <style>
  body {
    background-image: url('./assets/light.jpg');
    background-size: 100%;
  }

  .container {

    position: relative;
    height: 95vh;
  }

  .containercomp {
    border-radius: 20px;
    width: 1300px;
    height: 680px;
    position: absolute;
    right: 0;
    left: 0;
    margin: 0 auto;
    background-color: rgb(233, 248, 249);
    box-shadow: 0px 0px 50px 0px rgba(0, 0, 0, 0.5);
  }



  .listcomp {
    float: left;
    width: 14%;
  }

  .middlecomp {
    width: 70.5%;
    float: left;
  }

  .category{
    float: right;
    width: 14%;
    margin-right: 15px;
    background-color: skyblue;
  }
</style> -->