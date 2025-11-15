## R Programming

1. First Code

```

# Install packages (only if not already installed)
install.packages("palmerpenguins")  # use to install the library
install.packages("ggplot2")          # package used to scatter plot some data

# Load required packages (do this every session restart)
library(palmerpenguins)
library(ggplot2)

# Then create the plot
ggplot(data = penguins, aes(x = flipper_length_mm, y = body_mass_g)) +
  geom_point(aes(color = species))

```