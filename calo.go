package enronpeople

import "strings"

const usernamesCalo = `allen-p,arnold-j,arora-h,badeer-r,bailey-s,bass-e,baughman-d,beck-s,benson-r,blair-l,brawner-s,buy-r,campbell-l,carson-m,cash-m,causholli-m,corman-s,crandell-s,cuilla-m,dasovich-j,davis-d,dean-c,delainey-d,derrick-j,dickson-s,donoho-l,donohoe-t,dorland-c,ermis-f,farmer-d,fischer-m,forney-j,fossum-d,gang-l,gay-r,geaccone-t,germany-c,gilbertsmith-d,giron-d,griffith-j,grigsby-m,guzman-m,haedicke-m,hain-m,harris-s,hayslett-r,heard-m,hendrickson-s,hernandez-j,hodge-j,holst-k,horton-s,hyatt-k,hyvl-d,jones-t,kaminski-v,kean-s,keavey-p,keiser-k,king-j,kitchen-l,kuykendall-t,lavorato-j,lay-k,lenhart-m,lewis-a,linder-e,lokay-m,lokey-t,love-p,lucci-p,maggi-m,mann-k,martin-t,may-l,mccarty-d,mcconnell-m,mckay-b,mckay-j,mclaughlin-e,merriss-s,meyers-a,mims-thurston-p,motley-m,neal-s,nemec-g,panus-s,parks-j,pereira-s,perlingiere-d,phanis-s,pimenov-v,platter-p,presto-k,quenet-j,quigley-d,rapp-b,reitmeyer-j,richey-c,ring-a,ring-r,rodrique-r,rogers-b,ruscitti-k,sager-e,saibi-e,salisbury-h,sanchez-m,sanders-r,scholtes-d,schoolcraft-d,schwieger-j,scott-s,semperger-c,shackleton-s,shankman-j,shapiro-r,shively-h,skilling-j,slinger-r,smith-m,solberg-g,south-s,staab-t,stclair-c,steffes-j,stepenovitch-j,stokley-c,storey-g,sturm-f,swerzbin-m,symes-k,taylor-m,tholt-j,thomas-p,townsend-j,tycholiz-b,ward-k,watson-k,weldon-c,whalley-g,whalley-l,white-s,whitt-m,williams-j,williams-w3,wolfe-j,ybarbo-p,zipper-a,zufferli-j`

func UserNamesCalo() []string {
	return strings.Split(usernamesCalo, ",")
}
