import React from 'react'
import './style.css'

interface FormGroupProps {
  children: React.ReactNode
  direction?: 'row' | 'column'
  gap?: 'sm' | 'md' | 'lg'
}

export const FormGroup: React.FC<FormGroupProps> = ({
  children,
  direction = 'column',
  gap = 'md',
}) => {
  return (
    <div className={`form-group ${direction} gap-${gap}`}>
      {children}
    </div>
  )
}
