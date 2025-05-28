import './globals.css'

export const metadata = {
  title: 'Sistema Hospitalar CESUPA',
  description: 'Sistema de Gestão Hospitalar desenvolvido com Next.js',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="pt-BR">
      <body>
        <main>{children}</main>
      </body>
    </html>
  )
}
