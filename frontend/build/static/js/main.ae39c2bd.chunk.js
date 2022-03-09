(this["webpackJsonpdropbox-to-ipfs-frontend"]=this["webpackJsonpdropbox-to-ipfs-frontend"]||[]).push([[0],{41:function(e,t,c){},43:function(e,t,c){"use strict";c.r(t);var a=c(1),s=c(24),r=c.n(s),n=c(4),i=c(12),o=c(13),l=c(14),d=c(17),j=c(44),h=c(45),b=c(46),m=c(25),x=c.n(m),u=c(2),O=function(e){Object(l.a)(c,e);var t=Object(d.a)(c);function c(e){var a;return Object(i.a)(this,c),(a=t.call(this,e)).state={verification:x()()},a}return Object(o.a)(c,[{key:"render",value:function(){var e="https://www.dropbox.com/oauth2/authorize?response_type=code&client_id=".concat("c3hbpngaqu240bf","&redirect_uri=").concat("http://localhost:3200/api/dropbox/oauth_callback","&state=").concat(this.state.verification);return Object(u.jsx)(j.a,{children:Object(u.jsx)(h.a,{children:Object(u.jsx)(b.a,{children:Object(u.jsx)("a",{href:e,className:"btn btn-primary btn-lg mt-4",children:"Primary link"})})})})}}]),c}(a.Component),p=c(47),f=c(11),v=c(21),g=c(19),y=function(e){Object(l.a)(c,e);var t=Object(d.a)(c);function c(e){var a;return Object(i.a)(this,c),(a=t.call(this,e)).state={callbackUID:null,isError:!1},a}return Object(o.a)(c,[{key:"componentDidMount",value:function(){var e=this.props.match.params.token;null===e||void 0===e?this.setState({isError:!0}):this.setState({callbackUID:e})}},{key:"render",value:function(){var e=this.state,t=e.isError,c=e.callbackUID;return Object(u.jsx)(j.a,{children:Object(u.jsx)("div",{className:"jumbotron mt-5",children:t?Object(u.jsxs)(u.Fragment,{children:[Object(u.jsxs)("div",{className:"text-center",children:[Object(u.jsx)(v.a,{icon:g.b,size:"3x",style:{color:"orange"}}),Object(u.jsx)("h2",{className:"mt-4",children:"Something wrong during authentification flow"})]}),Object(u.jsx)(h.a,{className:"mt-4",children:Object(u.jsx)(b.a,{md:"12",children:Object(u.jsx)("div",{className:"card",children:Object(u.jsxs)("div",{className:"card-body text-center",children:["It seems an error occured when validating your authentification to your Dropbox account. Please try again or create an issue on our Github repository",Object(u.jsx)("br",{}),Object(u.jsx)("br",{}),Object(u.jsx)("a",{href:"https://github.com/FlorianRuen/Dropbox-To-IPFS-App",children:"https://github.com/FlorianRuen/Dropbox-To-IPFS-App"})]})})})}),Object(u.jsx)(h.a,{className:"text-center mt-4",children:Object(u.jsx)(b.a,{md:"12",children:Object(u.jsx)(p.a,{tag:f.b,to:"/",color:"warning",size:"3x",children:Object(u.jsx)("span",{className:"as--light",children:"Try authentification again"})})})})]}):Object(u.jsxs)(u.Fragment,{children:[Object(u.jsxs)("div",{className:"text-center",children:[Object(u.jsx)(v.a,{icon:g.a,size:"3x",style:{color:"green"}}),Object(u.jsx)("h2",{className:"mt-4",children:"Integration with your Dropbox account completed"})]}),Object(u.jsxs)(h.a,{className:"mt-4",children:[Object(u.jsx)(b.a,{md:"4",children:Object(u.jsxs)("div",{className:"card",children:[Object(u.jsx)("div",{className:"card-header font-weight-bold",children:"1. Copy / move files to app folder"}),Object(u.jsx)("div",{className:"card-body",children:"First step, copy all the files you want to migrate to IPFS in the folder named Applications/Send-To-IPFS"})]})}),Object(u.jsx)(b.a,{md:"4",children:Object(u.jsxs)("div",{className:"card",children:[Object(u.jsx)("div",{className:"card-header font-weight-bold",children:"2. Your files will be hosted on IPFS"}),Object(u.jsx)("div",{className:"card-body",children:"Our application will automatically detect all the files you add to send them over IPFS, without deleting the original"})]})}),Object(u.jsx)(b.a,{md:"4",children:Object(u.jsxs)("div",{className:"card",children:[Object(u.jsx)("div",{className:"card-header font-weight-bold",children:"3. Check migrated files and status"}),Object(u.jsx)("div",{className:"card-body",children:"Using the identifier below, you can access to the application to check all files that were migrated and their status"})]})})]}),Object(u.jsx)(h.a,{className:"mt-4",children:Object(u.jsx)(b.a,{md:"12",children:c})})]})})})}}]),c}(a.PureComponent),N=function(){return Object(u.jsx)("div",{className:"app-container d-flex flex-column",children:Object(u.jsx)("div",{children:Object(u.jsxs)(n.c,{children:[Object(u.jsx)(n.a,{exact:!0,path:"/",component:O}),Object(u.jsx)(n.a,{path:"/callback/:token?",component:y})]})})})};c(41),c(42);r.a.render(Object(u.jsx)(f.a,{children:Object(u.jsx)(N,{})}),document.getElementById("root"))}},[[43,1,2]]]);
//# sourceMappingURL=main.ae39c2bd.chunk.js.map