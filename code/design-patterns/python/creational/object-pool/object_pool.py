class Bullet:
    def init(self):
        self.position = None

    def move(self):
        # Move the bullet on the screen
        pass

    def draw(self):
        # Draw the bullet on the screen
        pass


class BulletPool:
    _instance = None
    maxBullets = 10  # Setting the maximum bullets allowed in the pool

    @staticmethod
    def getInstance():
        if BulletPool._instance is None:
            BulletPool._instance = BulletPool(BulletPool.maxBullets)
        return BulletPool._instance

    def __init__(self, size):
        if BulletPool._instance is not None:
            raise Exception("This class is a singleton!")
        else:
            # Initialize the pool with up to 'size' Bullets, capped by maxBullets
            self.pool = [Bullet() for _ in range(min(size, BulletPool.maxBullets))]

    def acquire(self):
        if self.pool:
            return self.pool.pop()
        else:
            return Bullet()  # Optionally, you might want to handle this differently

    def release(self, bullet: Bullet):
        if len(self.pool) < BulletPool.maxBullets:
            self.pool.append(bullet)
        else:
            # Optionally handle or log when the pool is at capacity and cannot accept more bullets
            pass
