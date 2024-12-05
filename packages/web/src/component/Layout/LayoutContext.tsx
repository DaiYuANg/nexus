import { createContext, ReactNode, useContext } from 'react';

type LayoutContextProp = {
  headerHeight: number;
  navbarWidth: number;
};

const LayoutContext = createContext<LayoutContextProp | undefined>(undefined);

type LayoutProviderProps = {
  children: ReactNode;
  headerHeight: number;
  navbarWidth: number;
};

const LayoutProvider = ({ children, headerHeight, navbarWidth }: LayoutProviderProps) => {
  return (
    <LayoutContext.Provider
      value={{
        headerHeight,
        navbarWidth,
      }}
    >
      {children}
    </LayoutContext.Provider>
  );
};

const useLayoutContext = () => {
  const context = useContext(LayoutContext);
  if (!context) {
    throw new Error('useChatContext must be used within a ProductProvider');
  }
  return context;
};

export { LayoutProvider, useLayoutContext };
