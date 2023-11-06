<template>
    <div>
        <div class="middlelist1">
            <VideoList />
        </div>
        <div class="middlelist2">
            <VideoDetailsright />
        </div>
        <!-- <div class="rightlist">
            <VideoCategory />
        </div> -->
    </div>
</template>


<script>

    import VideoList from '../components/VideoList.vue'
    import VideoDetailsright from '../components/VideoDetailsright.vue'
    // import VideoCategory from '../components/VideoCategory.vue'
    export default {
        name: 'ShowandCate',
        components: {
            VideoList,
            VideoDetailsright,
            // VideoCategory,
        }
    }
</script>


<style>
    .middlelist1 {
        height: 92vh;
        width: 75vw;
        float: left;
        /* background-color: #fff; */
    }

    .middlelist2 {
        /* background-color: #fff; */
    }

    .rightlist {
        height: 92vh;
        width: 5vw;
        float: left;
        /* background-color: #fff; */
    }
</style>