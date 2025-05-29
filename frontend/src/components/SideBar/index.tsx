'use client'

import './style.css'
import { usePathname } from 'next/navigation'

export const Sidebar = () => {
  const pathname = usePathname()

  const links = [
    { label: 'Agenda', icon: '📅', href: '/agenda' },
    { label: 'Pacientes', icon: '🧍', href: '/pacientes' },
  ]

  return (
    <aside className="sidebar">
      <div className="logo">código<span>med</span></div>
      <nav>
        {links.map(link => (
          <a key={link.href} href={link.href} className={pathname === link.href ? 'active' : ''}>
            <span className="icon">{link.icon}</span> {link.label}
          </a>
        ))}
      </nav>
    </aside>
  )
}
