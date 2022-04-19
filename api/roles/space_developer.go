package roles

//+kubebuilder:rbac:groups="",resources=secrets,verbs=patch;get;create
//+kubebuilder:rbac:groups="",resources=pods,verbs=list

//+kubebuilder:rbac:groups=workloads.cloudfoundry.org,resources=cfapps,verbs=get;create;patch;delete;list
//+kubebuilder:rbac:groups=workloads.cloudfoundry.org,resources=cfprocesses,verbs=get;create;patch;list
//+kubebuilder:rbac:groups=workloads.cloudfoundry.org,resources=cfpackages,verbs=get;create;patch;list
//+kubebuilder:rbac:groups=workloads.cloudfoundry.org,resources=cfbuilds,verbs=get;create;list

//+kubebuilder:rbac:groups=services.cloudfoundry.org,resources=cfserviceinstances,verbs=get;create;list;delete
//+kubebuilder:rbac:groups=services.cloudfoundry.org,resources=cfservicebindings,verbs=get;create;list;delete

//+kubebuilder:rbac:groups=networking.cloudfoundry.org,resources=cfroutes,verbs=get;list;create;delete;patch

//+kubebuilder:rbac:groups=kpack.io,resources=clusterbuilders,verbs=get
