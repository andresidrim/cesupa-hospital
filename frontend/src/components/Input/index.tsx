'use client'

import React from 'react'
import './style.css'

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string
  error?: string
}

export const Input: React.FC<InputProps> = ({ label, error, ...props }) => {
  return (
    <div className="input-group">
      {label && <label htmlFor={props.id || props.name}>{label}</label>}
      <input className={`input-field ${error ? 'input-error' : ''}`} {...props} />
      {error && <span className="input-message">{error}</span>}
    </div>
  )
}
