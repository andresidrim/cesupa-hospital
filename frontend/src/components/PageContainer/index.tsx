import './style.css'
import { Sidebar } from '@/components/SideBar';

export const PageContainer = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="layout">
      <Sidebar />
      <div className="content">
        {children}
      </div>
    </div>
  )
}
