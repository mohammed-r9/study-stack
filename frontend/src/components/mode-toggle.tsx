import { Moon, Sun } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useThemeStore } from '@/lib/store/theme'

type Props = {
  variant?: 'link' | 'default' | 'ghost'
}
export default function ModeToggle({ variant }: Props) {
  const toggleTheme = useThemeStore((state) => state.toggleTheme)
  return (
    <Button
      variant={variant || 'ghost'}
      size="icon"
      className="cursor-pointer"
      onClick={toggleTheme}
    >
      <Sun className="h-[1.2rem] w-[1.2rem] scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90" />
      <Moon className="absolute h-[1.2rem] w-[1.2rem] scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0" />
      <span className="sr-only">Toggle theme</span>
    </Button>
  )
}
