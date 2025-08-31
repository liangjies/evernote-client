import router from '@/router/index'
router.beforeEach((to,from,next)=>{
	if(to.meta&&to.meta.title){
	document.title = to.meta.title
	}
	next()
})