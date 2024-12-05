import { ComponentType, FC, lazy, LazyExoticComponent, ReactNode, Suspense } from 'react';
import { LoadingOverlay } from '@mantine/core';

type LazyLoadFunc = () => Promise<{ default: ComponentType<any> }>;

const lazyLoad = (importFunc: LazyLoadFunc): LazyExoticComponent<React.ComponentType<any>> =>
  lazy(() => importFunc().then((module) => ({ default: module.default })));

interface LazyComponentWrapperProps<T> {
  importFunc: LazyLoadFunc;
  fallback?: ReactNode; // 可选的加载中显示内容
  props: T;
}

// 使用 Suspense 包裹懒加载组件，展示 loading 或错误信息
const LazyComponentWrapper: FC<LazyComponentWrapperProps<any>> = ({
  importFunc,
  fallback = (
    <LoadingOverlay
      visible={true}
      zIndex={1000}
      overlayProps={{
        radius: 'sm',
        blur: 3,
      }}
    />
  ),
  props,
}) => {
  const LazyComponent = lazyLoad(importFunc);

  return (
    <Suspense fallback={fallback}>
      <LazyComponent {...props} />
    </Suspense>
  );
};

export { LazyComponentWrapper };
