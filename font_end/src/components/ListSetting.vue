<template>
  <div class="menu-wrapper" >
    <div class="content">
      <div class="menu-content">
        <div class="menu-list">
          <div class="menu-list-item" v-for="item in menuData" :key="item.id" @click="pushshow(item.id)">
            <i :class="item.iconname"></i>
            <!-- <div class="item-name">{{ item.menuName }}</div> -->
          </div>
          <!-- <div class="menu-list-item">
            <div class="block"></div>
            <i class="el-icon-caret-right"></i>
            <div class="item-name">{{ Darkorlight }}</div>
            <div class="btn">
              <input class="check-ipt" type="checkbox" id="checks" />
              <label
                class="check-lable"
                for="checks"
                @click="changeDark"
              ></label>
            </div>
          </div> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:'ListSetting',
  data() {
    return {
      Darkorlight:"夜间模式",
      menuData: [
        {id: 1,menuName: "首页",iconname: "el-icon-monitor",},
        {id: 2,menuName: "个人中心",iconname: "el-icon-s-custom",},
        {id: 3,menuName: "分类",iconname: "el-icon-s-grid",},
        // {id: 4,menuName: "浏览记录",iconname: "el-icon-s-order",},
        {id: 4,menuName: "设置",iconname: "el-icon-s-tools",},
        // {id: 6,menuName: "退出登录",iconname: "el-icon-error",},
      ],
    };
  },
  methods: {
    changeDark() {
      if(this.Darkorlight=='夜间模式') {
        this.Darkorlight='日间模式'
      }
      else{
        this.Darkorlight='夜间模式'
      }
    },
    pushshow(params){
      // console.log(params);
      this.$bus.$emit('showpage',params);
    }
  },
};
</script>


<style scoped>

.iconname {
  font-size: 20px;
  color: grey;
  vertical-align: middle;
}

.menu-wrapper {

  height: 92vh;
  width:25px;
  background-color: black;
  border-right: 1px solid white;
  /* background-color: rgb(233, 248, 249); */
  /* padding-left: 10px; */
  padding-right: 8px;
  padding-top: 1px;
  /* transition: 0.4s; */
}
/* 
.menu-wrapper::after{
    opacity: 0;
  } */

.menu-wrapper .menu-title {
  padding-bottom: 15px;
  box-sizing: border-box;
  border-bottom: 1px solid rgb(229, 233, 236);
  margin-bottom: 10px;
}

.menu-wrapper .menu-title .title-text {
  /* margin-left: 10px; */
  font-family: "BasketBall";
  font-size: 24px;
  vertical-align: middle;
  /* opacity: 0; */
  /* transition: 0.4s; */
  color: var(--theme-info-text-color);
}

.detail{
  font-size: 13px;
}

.menu-wrapper .content .person-info {
  margin-top: 20px;
  white-space: nowrap;
}

.menu-wrapper .content .person-info img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  border: 2px solid #fff;
  box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.2);
  vertical-align: middle;
}

.menu-wrapper .content .person-info .person-name {
  /* margin-left: 15px; */
  vertical-align: middle;
  /* transition: 0.4s; */
  overflow: hidden;
  display: inline-block;
}

i{
  font-size: 25px;
  color: grey;
}

.menu-list-item {
  flex: 1;
  width: 100%;
  height: 25px;
  position: relative;
  padding-top: 13px;
  padding-bottom: 15px;
  padding-left: 2px;
  white-space: nowrap;
  cursor: pointer;
}


.item-name {
  line-height: 50px;
  display: inline-block;
  margin-left: 15px;
  color: grey;
  font-weight: bold;
}

.menu-list-item:hover .item-name {
  color: black;
}

.menu-wrapper .content .menu-content .menu-list .menu-list-item .btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  display: inline-block;
  margin-left: 10px;
}
.menu-wrapper .content .menu-content .menu-list .menu-list-item .btn .check-ipt {
  display: none;
}

.menu-wrapper .content .menu-content .menu-list .menu-list-item .btn .check-lable {
  display: inline-block;
  width: 40px;
  height: 20px;
  border-radius: 20px;
  background: rgb(101, 58, 223);
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.menu-wrapper .content .menu-content .menu-list .menu-list-item .btn .check-lable::before {
  display: inline-block;
  content: "";
  width: 15px;
  height: 15px;
  border-radius: 50%;
  background: white;
  position: absolute;
  left: 3px;
  top: 50%;
  transform: translateY(-50%);
  transition: all 0.3s;
}

.menu-wrapper .content .menu-content .menu-list .menu-list-item .btn .check-ipt:checked + .check-lable::before {
  left: 20px;
  background: rgb(32, 30, 43);
}

.menu-wrapper .content .menu-content .menu-list .menu-list-item .btn .check-ipt:checked + .check-lable::after {
  width: 100%;
}

.menu-wrapper:hover .person-info .person-name {
  opacity: 1;
}


.menu-wrapper:hover .menu-content .menu-list .menu-list-item .btn {
  opacity: 1;
}

</style>
