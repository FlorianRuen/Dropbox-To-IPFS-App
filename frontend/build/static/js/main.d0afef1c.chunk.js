(this["webpackJsonpdropbox-to-ipfs-frontend"]=this["webpackJsonpdropbox-to-ipfs-frontend"]||[]).push([[0],{84:function(t,n,e){},85:function(t,n,e){"use strict";e.r(n);var o=e(0),r=e(39),c=e.n(r),i=e(2),a=e(65),s=e(40),u=e(41),l=e(42),p=e(66),h=e(88),d=e(45),b=e(89),f=e(62),x=e(10),j=e(43),v=e.n(j),k=e(44),m=e.n(k),O=e(7),y=function(t){Object(l.a)(e,t);var n=Object(p.a)(e);function e(t){var o;return Object(s.a)(this,e),(o=n.call(this,t)).state={isDropboxInit:!1,apiToken:"",verification:v()()},o}return Object(u.a)(e,[{key:"componentDidMount",value:function(){var t=this;h.a.addEventListener("url",(function(n){return t.handleLinkingUrl(n)}))}},{key:"componentWillUnmount",value:function(){var t=this;h.a.removeEventListener("url",(function(n){return t.handleLinkingUrl(n)}))}},{key:"handleLinkingUrl",value:function(t){var n=t.url.match(/(.*)/),e=Object(a.a)(n,2)[1],o=m()(e);this.state.verification===o.state?this.setState({isDropboxInit:!0,apiToken:o.access_token}):alert("Invalid checksum")}},{key:"loginWithDropbox",value:function(){var t="https://www.dropbox.com/oauth2/authorize?response_type=code&client_id=".concat("c3hbpngaqu240bf","&redirect_uri=").concat("http://localhost:3200/api/dropbox/oauth_callback","&state=").concat(this.state.verification);h.a.openURL(t).catch((function(t){return console.error("An error occurred",t)}))}},{key:"render",value:function(){var t=this,n=this.state.isDropboxInit?"Dropbox API token : "+this.state.apiToken:"You are not already connected to your Dropbox account";return Object(O.jsxs)(d.a,{style:g.container,children:[Object(O.jsx)(b.a,{title:"Connect my Dropbox account to IPFS",onPress:function(){return t.loginWithDropbox()}}),Object(O.jsx)(f.a,{style:g.instructions,children:n})]})}}]),e}(o.Component),g=x.a.create({container:{flex:1,justifyContent:"center",alignItems:"center",backgroundColor:"#F5FCFF"},instructions:{marginTop:32,textAlign:"center",color:"#333333",marginBottom:5}}),D=function(){return Object(O.jsx)("div",{className:"app-container d-flex flex-column",children:Object(O.jsx)("div",{children:Object(O.jsx)(i.c,{children:Object(O.jsx)(i.a,{exact:!0,path:"/",component:y})})})})},I=e(34);e(84);c.a.render(Object(O.jsx)(I.a,{children:Object(O.jsx)(D,{})}),document.getElementById("root"))}},[[85,1,2]]]);
//# sourceMappingURL=main.d0afef1c.chunk.js.map