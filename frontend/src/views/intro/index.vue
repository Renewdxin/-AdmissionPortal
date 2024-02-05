<template>
  <div class="intro">
    <div class="lab">
      <img src="@/assets/3g.jpg" alt="" />
      <div class="text">
        <div class="brief">
          <h1>
            XIYOU•MOBILE&nbsp;&nbsp;<em
              style="color: rgb(91, 194, 159); font-size: 22px"
              >/2011/</em
            >
          </h1>
        </div>
        <div class="detail">
          <p>西安邮电大学移动应用开发实验室，亦称西安邮电大学3G实验室。</p>
          <p>
            创建于2011年，隶属于西安邮电大学计算机学院，专注于移动应用开发。
          </p>
          <p>2012年被谷歌授予“Android人才培养示范基地”</p>
          <p>
            实验室位于西安邮电大学长安校区东区教学楼一层FZ155教室，西邮IT应用型人才实训中心。
          </p>
          <p>实验室不仅为同学们提供了丰富的软、硬件资源，包括：</p>
          <p>
            120m2配备空调、独立桌子、靠背椅、饮水机的干净、整洁、舒适的学习环境。
          </p>
          <p>
            公共图书角、免费学习资料、科学上网、共享打印设备、投影仪等便利的学习资源。
          </p>
          <p>
            学习路线、周报或圆桌分享、官网wiki的学习分享、定时的讲课、老师指导的独特学习方式。
          </p>
          <p>毕业学长学姐的工作经验指导、春招秋招的内推直推优势的就业资源。</p>
          <p>
            并且一直以“传、帮、带”的学习方式，秉承“用代码编织未来”的宗旨，每年以百分百的就业率，源源不断的为国内著名互联网公司输送优秀人才。目前由Android、iOS、Web及后台4个小组组成~
          </p>
        </div>
      </div>
    </div>
    <div class="groups">
      <div class="group">
        <img src="@/assets/web.jpg" alt="" />
        <div class="nav">前端</div>
        <div class="text">
          Web致力于基于HTML&CSS&<br />&Vue&React的Web前端开发，以及以Node.js为主的Web后台开发。如今已上线多个项目，
          其中包括《xiyoumobile-admin》，《xiyoumobile-site》，《3gMusic》，《Interest》，《WVUE-UI》，
          《西邮读书会》等。我们追求最前沿的前端技术，在“斩码”的途中，大家也是互帮互助，披荆斩棘，培养了深厚的友谊。
          小组内学习氛围浓厚，闲暇时光更是欢乐无比，能让你找到真正的归属感。性感表现代码，把不可能变成可能。
        </div>
      </div>
      <div class="group">
        <img src="@/assets/server.jpg" alt="" />
        <div class="nav">后台</div>
        <div class="text">
          后端开发，也可以称为服务器端开发，是对用户看不到的程序后端部分的开发，负责处理前端的请求，进行逻辑处理和数据交互，涉及到mysql、redis、
          后台语言及相关框架等相关知识。后端开发需要考虑底层业务逻辑的实现、数据的保存与读取、平台的稳定性和性能等。
        </div>
      </div>
      <div class="group">
        <img src="@/assets/android.jpg" alt="" />
        <div class="nav">安卓</div>
        <div class="text">
          Android致力于 Android 移动平台的 APP
          开发。组内活跃着学习的气息和交流的热情， 是 Google 大学合作部——Android
          人才培养示范基地。在这样的环境下，我们努力创造出自己的 APP，
          取得Google
          全国大学生移动互联网创新挑战赛金奖、银奖，爱奇艺全国高校技术大赛最佳创新奖等等。
          我们，因共同热爱而聚集，为开发梦想而诞生。
        </div>
      </div>
      <div class="group">
        <img src="@/assets/ios.jpg" alt="" />
        <div class="nav">iOS</div>
        <div class="text">
          iOS致力基于iOS系统的移动应用开发，主要使用Objective-C语言进行开发。有着先进的iOS开发设备，
          积极活跃的开发团队，团队成员互相合作，共同进步，积极参加各类竞赛。也正因如此，近几年来我们都取得
          了非常可观的成绩，如今已经有玩转西邮，西邮成绩，西邮图书馆，近健步行等多个APP上线，除此之外我们的
          就业率更是达到了百分之百，并且都是就职于腾讯，百度，爱奇艺，搜狐等大型互联网公司。怀揣着对iOS开发
          的无限热情，我们将继续努力，共同进步。
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
  text-decoration: none;
}
/* 电脑平板设备 */
@media (min-width: 640px) {
  .intro {
    background: url('@/assets/bkg27.jpg') no-repeat;
    background-size: 100% 100%;
    margin-top: -53px;
    margin-left: -8px;
  }
  .lab,
  .groups {
    width: 84vw;
    margin: 7vh auto;
    box-sizing: content-box;
    border: 4px solid rgb(121, 115, 115);
    background-color: rgba(255, 255, 255, 0.4);
  }
  .lab {
    position: relative;
    margin-bottom: 0vh;
    height: 400px;
    border-top-left-radius: 90px;
    img {
      position: absolute;
      left: 0px;
      width: 180px;
      border-radius: 50%;
      margin-right: 30px;
    }
    .text {
      width: 60vw;
      height: 100%;
      position: absolute;
      right: 0px;
      color: rgb(47, 57, 56); //介绍小字的颜色
      top: 0px;
      .brief h1 {
        margin-top: 20px;
        margin-bottom: 14px;
      }
      .detail {
        // background-color: pink;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        overflow: hidden;
        -webkit-line-clamp: 15;
        text-overflow: ellipsis;
      }
    }
  }
  .groups {
    margin-top: 4vh;
    margin-bottom: 0vh;
    display: flex;
    justify-content: space-evenly;
    padding-top: 20px;
    padding-bottom: 20px;
    height: 860px;
    .group {
      box-sizing: content-box;
      border: 4px solid rgb(138, 138, 138);
      position: relative;
      width: 18%;
      height: 800px;
      // background-color: pink;
      padding-left: 20px;
      padding-right: 10px;
      .nav {
        width: 100px;
        height: 4vw;
        border-radius: 10px;
        line-height: 3.5vw;
        text-align: center;
        color: azure;
        font-size: 2.3vw;
        position: absolute;
        left: 50%;
        margin-left: -50px;
        background-color: rgb(95, 95, 95);
      }
      &:nth-child(even) {
        border-bottom-left-radius: 15.12vw; //84vw*18%=15.12vw
        border-bottom-right-radius: 15.12vw;
        img {
          bottom: 0px;
        }
        .text {
          width: 85%;
          position: absolute;
          top: 6vw;
          color: rgb(47, 57, 56); //介绍小字的颜色
        }
        .nav {
          top: 1vw;
        }
      }
      &:nth-child(odd) {
        border-top-left-radius: 15.12vw;
        border-top-right-radius: 15.12vw;
        img {
          top: 0px;
        }
        .text {
          width: 85%;
          position: absolute;
          color: rgb(47, 57, 56); //介绍小字的颜色
          top: 21vw;
        }
        .nav {
          top: 16vw;
        }
      }
      img {
        width: 15.12vw;
        border-radius: 50%;
        position: absolute;
        left: 50%;
        margin-left: -7.56vw;
      }
      .text {
        display: -webkit-box;
        -webkit-box-orient: vertical;
        overflow: hidden;
        -webkit-line-clamp: 30;
        text-overflow: ellipsis;
        // background-color: palegreen;
      }
    }
  }
}
</style>
