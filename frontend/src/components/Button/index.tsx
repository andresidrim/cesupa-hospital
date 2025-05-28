'use client'

import React from 'react'
import './style.css'

type Variant = 'default' | 'outline' | 'ghost' | 'destructive'

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: Variant
}

export const Button: React.FC<ButtonProps> = ({ variant = 'default', className = '', children, ...props }) => {
  return (
    <button
      className={`button ${variant} ${className}`}
      {...props}
    >
      {children}
    </button>
  )
}
