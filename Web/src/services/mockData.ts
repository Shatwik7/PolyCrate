import { Asset, User, Comment } from '@/types/api';

const defaulturl='https://incidentscreenshorts.s3.ap-south-1.amazonaws.com/free_1972_datsun_240k_gt.glb?response-content-disposition=inline&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEDEaCmFwLXNvdXRoLTEiRjBEAiBZIKNKNPkxhknbt%2FhfBEAg8D9KKwNXgShGfR6h0FQxjQIgCpUow%2FGOY4sCj0M1q%2BKzIiDzEaf2bxOOknkUBS02w44quQMIahAAGgw5MDgwMjc0MDMyMDEiDCt1slRdY7A4UhOMriqWA7u39Eexu9qOXu4cSw3fouH2kNbMNwRfLxiqmpruU53i29r6cvhBhYI6mr3nk5qpqYkc9gK2mJ%2FPZhbfS3cjJ1rTjAR2tgFXzYKjiRcaZSgqVdfO1tuQEJrw5LkSR6jcHzh1WyOXzLXOkkfZKn%2Bl7j5X91iK4sHq0dKTlR8p%2F26ha9d%2FPL45mlX1Ddm6OILvwRQtyotWNNjgt0yh%2FtwqtRNsf05Hhmq5C78f%2BZtQxcsvv1sEnI0mKWL5bM11MvYOEfKGWGJdvlbMPrpigYvENIU6m5kEnczx%2FWu27rQ0r8OoOzjaQ5gM2f1RpvFTNjEHcG6dTPdEfdcqQFKH0En%2F7umk8xiZe8dmtKZtZU12fb1KU8BYekQy20VGsx6K7bnIqnEGu7ZzP1ECxmlAKD4Rq0nzk7HStlfIvlPPIV9fP%2Ft0pLsmygOWvZWpJE1eA9OquajEnWyMzv7FsD6JQRlHV%2BYdBuZla1In6tWV9b7af6%2BOq2PYk1uaKSblXDWGTqys1GBI52q4fczheiY9%2BiO%2BkBucIsLoWvYwp%2BTJxAY63wLBBQhRha5PYLswjYhP7H2MbFtNrZ4mJGa1ErrMsab%2F2p2pNS5ZIAKeHpepN4C26yOXdOuOKCW%2FMHJQRELoGg0HiWgeQUKqp7hJclXWDxcXO%2Bu5x%2Fymr%2BGe62D32eMrCs4tLiymmLogzLxxxrqkdaDI0ZIH8N%2F67Yy9%2BZtfmk3b%2BOWqPZlePvXqXjTMVguQx1oT8hW3KyBOChqk%2B5tpN4%2F%2BV3rOhmSGXZHlLWMeZLci2CvxjMYmS7YVVvFG%2FGfwzUiCpB2vL4ZzyYmtl%2Bf37ZUisfqF%2F3HfzFiw5ohPApsZeYH2Nxd3ZrQjseMBwOBbzEpZwoa3GbAeMNKDqyXaDDz0n25gTyBEySJV41obfNFBXrv3ZbVHhu4g9dLGN1O%2BdllmwsEgXHG2E1%2Fx30L%2B5m%2BVzfvw5ZXjRYTv%2F05EQ%2BqVAI9GvB0FrMY3CU0EJ6GnxVFljvyUYyGobiYLTYtWpJI%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA5G2VGTPA345JYW2V%2F20250806%2Fap-south-1%2Fs3%2Faws4_request&X-Amz-Date=20250806T003627Z&X-Amz-Expires=43200&X-Amz-SignedHeaders=host&X-Amz-Signature=a6de95c1cafb9355bac069fc3dcc4a3fa8cd645282b67ecfd3d426f2917d7f34';

export const mockUsers: User[] = [
  {
    id: '1',
    username: 'alex_designer',
    email: 'alex@example.com',
    full_name: 'Alex Rodriguez',
    profile_picture_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex',
    bio: '3D Artist & Designer passionate about creating immersive experiences',
    created_at: '2024-01-15T10:00:00Z',
    updated_at: '2024-03-01T15:30:00Z'
  },
  {
    id: '2',
    username: 'maya_creator',
    email: 'maya@example.com',
    full_name: 'Maya Chen',
    profile_picture_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Maya',
    bio: 'Game developer and 3D model enthusiast. Love creating stylized characters.',
    created_at: '2024-02-10T14:20:00Z',
    updated_at: '2024-03-02T12:45:00Z'
  },
  {
    id: '3',
    username: 'retro_models',
    email: 'retro@example.com',
    full_name: 'Sam Mitchell',
    profile_picture_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sam',
    bio: 'Specializing in retro and vintage 3D models. Bringing the past to digital life.',
    created_at: '2024-01-20T09:15:00Z',
    updated_at: '2024-02-28T16:20:00Z'
  }
];

export const mockAssets: Asset[] = [
  {
    id: '1',
    creator_id: '1',
    file_name: 'modern_chair.glb',
    file_url: defaulturl,
    file_format: 'glb',
    preview_url: 'https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=400',
    thumbnail_url: 'https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=200',
    downloads: 1247,
    likes: 89,
    description: 'A sleek, modern chair perfect for contemporary interiors. Optimized for real-time rendering.',
    created_at: '2024-02-15T10:30:00Z',
    updated_at: '2024-02-15T10:30:00Z',
    is_public: true,
    creator: mockUsers[0],
    tags: ['furniture', 'modern', 'chair', 'interior'],
    metadata: { 'polygons': '2500', 'textures': '4K' }
  },
  {
    id: '2',
    creator_id: '2',
    file_name: 'cyberpunk_character.glb',
    file_url: defaulturl,
    file_format: 'glb',
    preview_url: 'https://images.unsplash.com/photo-1518709268805-4e9042af2176?w=400',
    thumbnail_url: 'https://images.unsplash.com/photo-1518709268805-4e9042af2176?w=200',
    downloads: 2156,
    likes: 234,
    description: 'Futuristic cyberpunk character with detailed textures and rigging. Ready for animation.',
    created_at: '2024-02-20T14:45:00Z',
    updated_at: '2024-02-20T14:45:00Z',
    is_public: true,
    creator: mockUsers[1],
    tags: ['character', 'cyberpunk', 'sci-fi', 'rigged'],
    metadata: { 'polygons': '8500', 'textures': '2K', 'rigged': 'yes' }
  },
  {
    id: '3',
    creator_id: '3',
    file_name: 'vintage_car.obj',
    file_url: defaulturl,
    file_format: 'obj',
    preview_url: 'https://images.unsplash.com/photo-1503376780353-7e6692767b70?w=400',
    thumbnail_url: 'https://images.unsplash.com/photo-1503376780353-7e6692767b70?w=200',
    downloads: 987,
    likes: 156,
    description: 'Classic 1950s vintage car model with authentic details and weathered textures.',
    created_at: '2024-02-10T16:20:00Z',
    updated_at: '2024-02-10T16:20:00Z',
    is_public: true,
    creator: mockUsers[2],
    tags: ['vehicle', 'vintage', 'car', 'retro'],
    metadata: { 'polygons': '12000', 'textures': '4K' }
  },
  {
    id: '4',
    creator_id: '1',
    file_name: 'space_station.gltf',
    file_url: defaulturl,
    file_format: 'gltf',
    preview_url: 'https://images.unsplash.com/photo-1446776653964-20c1d3a81b06?w=400',
    thumbnail_url: 'https://images.unsplash.com/photo-1446776653964-20c1d3a81b06?w=200',
    downloads: 1876,
    likes: 298,
    description: 'Massive space station with intricate details and modular components. Perfect for sci-fi scenes.',
    created_at: '2024-01-25T11:15:00Z',
    updated_at: '2024-01-25T11:15:00Z',
    is_public: true,
    creator: mockUsers[0],
    tags: ['sci-fi', 'space', 'station', 'architecture'],
    metadata: { 'polygons': '45000', 'textures': '8K', 'animations': 'rotating_parts' }
  },
  {
    id: '5',
    creator_id: '2',
    file_name: 'magical_sword.glb',
    file_url: defaulturl,
    file_format: 'glb',
    preview_url: 'https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400',
    thumbnail_url: 'https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=200',
    downloads: 3421,
    likes: 567,
    description: 'Enchanted sword with glowing runes and particle effects. Game-ready with PBR materials.',
    created_at: '2024-03-01T09:30:00Z',
    updated_at: '2024-03-01T09:30:00Z',
    is_public: true,
    creator: mockUsers[1],
    tags: ['weapon', 'fantasy', 'sword', 'game-ready'],
    metadata: { 'polygons': '3200', 'textures': '2K', 'effects': 'particle_glow' }
  },
  {
    id: '6',
    creator_id: '3',
    file_name: 'cozy_cottage.glb',
    file_url: defaulturl,
    file_format: 'glb',
    preview_url: 'https://images.unsplash.com/photo-1518780664697-55e3ad937233?w=400',
    thumbnail_url: 'https://images.unsplash.com/photo-1518780664697-55e3ad937233?w=200',
    downloads: 892,
    likes: 143,
    description: 'Charming cottage house with detailed interiors and cozy atmosphere.',
    created_at: '2024-02-28T13:45:00Z',
    updated_at: '2024-02-28T13:45:00Z',
    is_public: true,
    creator: mockUsers[2],
    tags: ['architecture', 'house', 'cottage', 'cozy'],
    metadata: { 'polygons': '15000', 'textures': '4K', 'interiors': 'furnished' }
  }
];

export const mockComments: Comment[] = [
  {
    id: '1',
    asset_id: '1',
    user_id: '2',
    content: 'Amazing work! The topology is perfect for real-time rendering.',
    created_at: '2024-02-16T10:15:00Z',
    user: mockUsers[1]
  },
  {
    id: '2',
    asset_id: '1',
    user_id: '3',
    content: 'Could you share the Blender file? Would love to see the material setup.',
    created_at: '2024-02-17T14:30:00Z',
    user: mockUsers[2]
  },
  {
    id: '3',
    asset_id: '2',
    user_id: '1',
    content: 'The character design is incredible! Perfect for my cyberpunk game.',
    created_at: '2024-02-21T09:45:00Z',
    user: mockUsers[0]
  }
];

export const categories = [
  'All',
  'Characters',
  'Vehicles',
  'Architecture',
  'Furniture',
  'Weapons',
  'Nature',
  'Sci-Fi',
  'Fantasy',
  'Abstract'
];

export const popularTags = [
  'game-ready',
  'pbr',
  'low-poly',
  'high-poly',
  'rigged',
  'animated',
  'textured',
  'modern',
  'vintage',
  'fantasy',
  'sci-fi',
  'realistic',
  'stylized'
];