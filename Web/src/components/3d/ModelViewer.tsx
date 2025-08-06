import { Suspense, useRef } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { OrbitControls, PerspectiveCamera, Environment, ContactShadows } from '@react-three/drei';
import { Loader2 } from 'lucide-react';
import * as THREE from 'three';

interface ModelViewerProps {
  modelUrl?: string;
  className?: string;
  autoRotate?: boolean;
}

// Simple 3D shape component for demonstration
const DemoModel = ({ autoRotate = false }: { autoRotate?: boolean }) => {
  const meshRef = useRef<THREE.Mesh>(null);

  useFrame(() => {
    if (meshRef.current && autoRotate) {
      meshRef.current.rotation.y += 0.01;
    }
  });

  return (
    <group>
      {/* Main object */}
      <mesh ref={meshRef} position={[0, 0, 0]}>
        <torusKnotGeometry args={[1, 0.3, 128, 32]} />
        <meshStandardMaterial
          color="#8b5cf6"
          metalness={0.7}
          roughness={0.2}
          envMapIntensity={1}
        />
      </mesh>
      
      {/* Secondary objects */}
      <mesh position={[-2, 0, -1]}>
        <boxGeometry args={[0.8, 0.8, 0.8]} />
        <meshStandardMaterial
          color="#06b6d4"
          metalness={0.5}
          roughness={0.3}
        />
      </mesh>
      
      <mesh position={[2, 0, -1]}>
        <sphereGeometry args={[0.5, 32, 32]} />
        <meshStandardMaterial
          color="#f59e0b"
          metalness={0.6}
          roughness={0.2}
        />
      </mesh>
    </group>
  );
};

const LoadingSpinner = () => (
  <div className="absolute inset-0 flex items-center justify-center bg-background/80 backdrop-blur-sm">
    <div className="flex flex-col items-center gap-2">
      <Loader2 className="h-8 w-8 animate-spin text-primary" />
      <span className="text-sm text-muted-foreground">Loading 3D model...</span>
    </div>
  </div>
);

export const ModelViewer = ({ 
  modelUrl, 
  className = "w-full h-96", 
  autoRotate = true 
}: ModelViewerProps) => {
  return (
    <div className={`relative ${className} bg-gradient-to-br from-background to-muted rounded-lg overflow-hidden border border-border`}>
      <Canvas
        shadows
        camera={{ position: [0, 0, 5], fov: 50 }}
        gl={{ antialias: true, alpha: true }}
      >
        <Suspense fallback={null}>
          {/* Lighting */}
          <ambientLight intensity={0.6} />
          <directionalLight
            position={[5, 5, 5]}
            intensity={1}
            castShadow
            shadow-mapSize-width={2048}
            shadow-mapSize-height={2048}
          />
          <pointLight position={[-5, 5, 5]} intensity={0.5} color="#8b5cf6" />
          <pointLight position={[5, -5, -5]} intensity={0.3} color="#06b6d4" />

          {/* Environment */}
          <Environment preset="city" />
          
          {/* Model */}
          <DemoModel autoRotate={autoRotate} />
          
          {/* Ground shadow */}
          <ContactShadows
            position={[0, -2, 0]}
            opacity={0.4}
            scale={10}
            blur={2}
            far={4}
          />

          {/* Camera controls */}
          <OrbitControls
            enablePan={true}
            enableZoom={true}
            enableRotate={true}
            minDistance={3}
            maxDistance={10}
            autoRotate={autoRotate}
            autoRotateSpeed={1}
          />
        </Suspense>
      </Canvas>
      
      <Suspense fallback={<LoadingSpinner />}>
        {/* This would show loading state if we had actual model loading */}
      </Suspense>

      {/* Controls hint */}
      <div className="absolute bottom-2 left-2 bg-background/80 backdrop-blur-sm rounded px-2 py-1 text-xs text-muted-foreground">
        Click and drag to rotate • Scroll to zoom
      </div>
    </div>
  );
};