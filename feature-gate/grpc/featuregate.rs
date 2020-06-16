#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Feature {
    #[prost(string, tag = "1")]
    pub name: std::string::String,
    #[prost(bool, tag = "2")]
    pub open: bool,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FeatureName {
    #[prost(string, tag = "1")]
    pub name: std::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FeatureAddition {
    #[prost(string, tag = "1")]
    pub feature_name: std::string::String,
    #[prost(sfixed64, tag = "2")]
    pub guild_id: i64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FeatureRemoval {
    #[prost(string, tag = "1")]
    pub feature_name: std::string::String,
    #[prost(sfixed64, tag = "2")]
    pub guild_id: i64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GuildFeature {
    #[prost(string, tag = "1")]
    pub feature_name: std::string::String,
    #[prost(sfixed64, tag = "2")]
    pub guild_id: i64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FeatureList {}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreationResult {
    #[prost(bool, tag = "1")]
    pub success: bool,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OpennessResult {
    #[prost(bool, tag = "1")]
    pub open: bool,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddResult {
    #[prost(bool, tag = "1")]
    pub success: bool,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RemoveResult {
    #[prost(bool, tag = "1")]
    pub success: bool,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct FeatureResult {
    #[prost(bool, tag = "1")]
    pub success: bool,
}
#[doc = r" Generated server implementations."]
pub mod feature_gate_server {
    #![allow(unused_variables, dead_code, missing_docs)]
    use tonic::codegen::*;
    #[doc = "Generated trait containing gRPC methods that should be implemented for use with FeatureGateServer."]
    #[async_trait]
    pub trait FeatureGate: Send + Sync + 'static {
        async fn create_feature(
            &self,
            request: tonic::Request<super::Feature>,
        ) -> Result<tonic::Response<super::CreationResult>, tonic::Status>;
        async fn check_openness(
            &self,
            request: tonic::Request<super::FeatureName>,
        ) -> Result<tonic::Response<super::OpennessResult>, tonic::Status>;
        async fn add_guild_feature(
            &self,
            request: tonic::Request<super::FeatureAddition>,
        ) -> Result<tonic::Response<super::AddResult>, tonic::Status>;
        async fn remove_guild_feature(
            &self,
            request: tonic::Request<super::FeatureRemoval>,
        ) -> Result<tonic::Response<super::RemoveResult>, tonic::Status>;
        async fn check_guild_feature(
            &self,
            request: tonic::Request<super::GuildFeature>,
        ) -> Result<tonic::Response<super::FeatureResult>, tonic::Status>;
        #[doc = "Server streaming response type for the GetFeatures method."]
        type GetFeaturesStream: Stream<Item = Result<super::Feature, tonic::Status>>
            + Send
            + Sync
            + 'static;
        async fn get_features(
            &self,
            request: tonic::Request<super::FeatureList>,
        ) -> Result<tonic::Response<Self::GetFeaturesStream>, tonic::Status>;
    }
    #[derive(Debug)]
    #[doc(hidden)]
    pub struct FeatureGateServer<T: FeatureGate> {
        inner: _Inner<T>,
    }
    struct _Inner<T>(Arc<T>, Option<tonic::Interceptor>);
    impl<T: FeatureGate> FeatureGateServer<T> {
        pub fn new(inner: T) -> Self {
            let inner = Arc::new(inner);
            let inner = _Inner(inner, None);
            Self { inner }
        }
        pub fn with_interceptor(inner: T, interceptor: impl Into<tonic::Interceptor>) -> Self {
            let inner = Arc::new(inner);
            let inner = _Inner(inner, Some(interceptor.into()));
            Self { inner }
        }
    }
    impl<T, B> Service<http::Request<B>> for FeatureGateServer<T>
    where
        T: FeatureGate,
        B: HttpBody + Send + Sync + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = Never;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(&mut self, _cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/featuregate.FeatureGate/CreateFeature" => {
                    #[allow(non_camel_case_types)]
                    struct CreateFeatureSvc<T: FeatureGate>(pub Arc<T>);
                    impl<T: FeatureGate> tonic::server::UnaryService<super::Feature> for CreateFeatureSvc<T> {
                        type Response = super::CreationResult;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::Feature>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { inner.create_feature(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = CreateFeatureSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/featuregate.FeatureGate/CheckOpenness" => {
                    #[allow(non_camel_case_types)]
                    struct CheckOpennessSvc<T: FeatureGate>(pub Arc<T>);
                    impl<T: FeatureGate> tonic::server::UnaryService<super::FeatureName> for CheckOpennessSvc<T> {
                        type Response = super::OpennessResult;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::FeatureName>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { inner.check_openness(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = CheckOpennessSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/featuregate.FeatureGate/AddGuildFeature" => {
                    #[allow(non_camel_case_types)]
                    struct AddGuildFeatureSvc<T: FeatureGate>(pub Arc<T>);
                    impl<T: FeatureGate> tonic::server::UnaryService<super::FeatureAddition> for AddGuildFeatureSvc<T> {
                        type Response = super::AddResult;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::FeatureAddition>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { inner.add_guild_feature(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = AddGuildFeatureSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/featuregate.FeatureGate/RemoveGuildFeature" => {
                    #[allow(non_camel_case_types)]
                    struct RemoveGuildFeatureSvc<T: FeatureGate>(pub Arc<T>);
                    impl<T: FeatureGate> tonic::server::UnaryService<super::FeatureRemoval>
                        for RemoveGuildFeatureSvc<T>
                    {
                        type Response = super::RemoveResult;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::FeatureRemoval>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { inner.remove_guild_feature(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = RemoveGuildFeatureSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/featuregate.FeatureGate/CheckGuildFeature" => {
                    #[allow(non_camel_case_types)]
                    struct CheckGuildFeatureSvc<T: FeatureGate>(pub Arc<T>);
                    impl<T: FeatureGate> tonic::server::UnaryService<super::GuildFeature> for CheckGuildFeatureSvc<T> {
                        type Response = super::FeatureResult;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GuildFeature>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { inner.check_guild_feature(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = CheckGuildFeatureSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/featuregate.FeatureGate/GetFeatures" => {
                    #[allow(non_camel_case_types)]
                    struct GetFeaturesSvc<T: FeatureGate>(pub Arc<T>);
                    impl<T: FeatureGate> tonic::server::ServerStreamingService<super::FeatureList>
                        for GetFeaturesSvc<T>
                    {
                        type Response = super::Feature;
                        type ResponseStream = T::GetFeaturesStream;
                        type Future =
                            BoxFuture<tonic::Response<Self::ResponseStream>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::FeatureList>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { inner.get_features(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1;
                        let inner = inner.0;
                        let method = GetFeaturesSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.server_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => Box::pin(async move {
                    Ok(http::Response::builder()
                        .status(200)
                        .header("grpc-status", "12")
                        .body(tonic::body::BoxBody::empty())
                        .unwrap())
                }),
            }
        }
    }
    impl<T: FeatureGate> Clone for FeatureGateServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self { inner }
        }
    }
    impl<T: FeatureGate> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(self.0.clone(), self.1.clone())
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: FeatureGate> tonic::transport::NamedService for FeatureGateServer<T> {
        const NAME: &'static str = "featuregate.FeatureGate";
    }
}