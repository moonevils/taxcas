webpackJsonp([7],{Aqap:function(t,e){var n={},a=!1,o=!1,i=null,l=null;n.install=function(t,e){var n={defaultType:"bottom",duration:"2500",wordWrap:!1};for(var r in e)n[r]=e[r];t.prototype.$toast=function(e,o){var l=o||n.defaultType,r='<div v-show="show" :class="type" class="lx-toast '+(n.wordWrap?"lx-word-wrap":"")+'" '+(n.width?'style="width: '+n.width+'"':"")+">{{tip}}</div>";if(!a){if(!i){var d=t.extend({data:function(){return{show:a,tip:e,type:"lx-toast-"+l}},template:r}),s=(i=new d).$mount().$el;document.body.appendChild(s)}i.type="lx-toast-"+l,i.tip=e,i.show=a=!0,setTimeout(function(){i.show=a=!1},n.duration)}},["bottom","center","top"].forEach(function(e){t.prototype.$toast[e]=function(n){return t.prototype.$toast(n,e)}}),t.prototype.$loading=function(e,n){if("close"==n)l.show=o=!1;else{if(o)return;var a=t.extend({data:function(){return{show:o}},template:'<div v-show="show" class="lx-load-mark"><div class="lx-load-box"><div class="lx-loading"><div class="loading_leaf loading_leaf_0"></div><div class="loading_leaf loading_leaf_1"></div><div class="loading_leaf loading_leaf_2"></div><div class="loading_leaf loading_leaf_3"></div><div class="loading_leaf loading_leaf_4"></div><div class="loading_leaf loading_leaf_5"></div><div class="loading_leaf loading_leaf_6"></div><div class="loading_leaf loading_leaf_7"></div><div class="loading_leaf loading_leaf_8"></div><div class="loading_leaf loading_leaf_9"></div><div class="loading_leaf loading_leaf_10"></div><div class="loading_leaf loading_leaf_11"></div></div><div class="lx-load-content">'+e+"</div></div></div>"}),i=(l=new a).$mount().$el;document.body.appendChild(i),l.show=o=!0}},["open","close"].forEach(function(e){t.prototype.$loading[e]=function(n){return t.prototype.$loading(n,e)}})},t.exports=n},DL0H:function(t,e){},"LsP+":function(t,e){},NHnr:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var a=n("IvJb"),o={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{attrs:{id:"app"}},[e("router-view")],1)},staticRenderFns:[]};var i=n("C7Lr")({name:"App"},o,!1,function(t){n("LsP+")},null,null).exports,l=n("zO6J");a.a.use(l.a);var r=new l.a({routes:[{path:"/",name:"Home",component:function(t){return n.e(2).then(function(){var e=[n("HXef")];t.apply(null,e)}.bind(this)).catch(n.oe)}},{path:"/setform",name:"SetForm",component:function(t){return n.e(0).then(function(){var e=[n("lPi8")];t.apply(null,e)}.bind(this)).catch(n.oe)}},{path:"/order",name:"Order",component:function(t){return n.e(1).then(function(){var e=[n("rMuZ")];t.apply(null,e)}.bind(this)).catch(n.oe)}},{path:"/wait",name:"Wait",component:function(t){return n.e(3).then(function(){var e=[n("FCA4")];t.apply(null,e)}.bind(this)).catch(n.oe)}},{path:"/refund",name:"Refund",component:function(t){return n.e(5).then(function(){var e=[n("VEvN")];t.apply(null,e)}.bind(this)).catch(n.oe)}},{path:"/lookcert",name:"LookCert",component:function(t){return n.e(4).then(function(){var e=[n("crQ0")];t.apply(null,e)}.bind(this)).catch(n.oe)}}]}),d=n("aozt"),s=n.n(d),c=(n("D0oh"),n("ymEw")),u=n.n(c),p=(n("hTRY"),n("DL0H"),n("rVsN")),f=n.n(p);n("mi2a").polyfill();var v=s.a.create({baseURL:"https://tax.caishuidai.com",timeout:1e4});v.interceptors.request.use(function(t){return localStorage.token&&(t.headers.Authorization="Bearer "+localStorage.token),t},function(t){return alert("错误的传参","fail"),f.a.reject(t)}),v.interceptors.response.use(function(t){return t.headers.authorization&&(localStorage.token=t.headers.authorization),t},function(t){if(401===t.response.status||403===t.response.status);else if(500===t.response.status)return f.a.reject(t.response.data);return f.a.reject(t.response.data)});var h={install:function(t){Object.defineProperty(t.prototype,"$axios",{value:v})}},m=(n("XkZX"),n("Aqap")),_=n.n(m),g=n("efG5"),y=n.n(g);a.a.use(h),a.a.use(_.a),a.a.component("v-distpicker",u.a),a.a.config.productionTip=!1,a.a.prototype.$wx=y.a,a.a.prototype.baseUrl="https://tax.caishuidai.com",new a.a({el:"#app",router:r,components:{App:i},template:"<App/>"})},XkZX:function(t,e){},hTRY:function(t,e){let n=document.documentElement.clientWidth||document.body.clientWidth,a=document.getElementsByTagName("html")[0];n>750&&(n=750),a.style.fontSize=n/10+"px",window.addEventListener("resize",function(t){(n=document.documentElement.clientWidth||document.body.clientWidth)>750&&(n=750),a.style.fontSize=n/10+"px"})}},["NHnr"]);
//# sourceMappingURL=app.5035a65feeb176fb9c06.js.map