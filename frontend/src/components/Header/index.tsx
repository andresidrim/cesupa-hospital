import './style.css'

interface HeaderProps {
  title: string
  actionLabel?: string
  onActionClick?: () => void
}

export const Header: React.FC<HeaderProps> = ({ title, actionLabel, onActionClick }) => {
  return (
    <header className="page-header">
      <h1>{title}</h1>
      {actionLabel && (
        <button onClick={onActionClick} className="action-button">
          {actionLabel}
        </button>
      )}
    </header>
  )
}
