from object_pool import BulletPool

bullet_pool = BulletPool.getInstance()

# Simulate acquiring a bullet to fire
bullet1 = bullet_pool.acquire()
print(f"Bullet acquired and fired at position {bullet1.position}")

# Do something with the bullet, e.g., move it in a game loop
bullet1.position += 10  # Simulating the bullet moving
print(f"Bullet moved to position {bullet1.position}")

# Now, suppose the bullet hits a target or goes off screen
bullet_pool.release(bullet1)
print(f"Bullet released back to pool from position {bullet1.position}")

# Acquire another bullet
bullet2 = bullet_pool.acquire()
print(f"Bullet acquired and fired at position {bullet2.position}")

# Move this bullet
bullet2.position += 5
print(f"Bullet moved to position {bullet2.position}")

# Release the second bullet
bullet_pool.release(bullet2)
print(f"Bullet released back to pool from position {bullet2.position}")
