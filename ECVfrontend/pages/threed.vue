<template>
  <div ref="demo1"></div>
</template>

<script>
import * as Three from 'three'
//import dat from 'dat.gui'
export default {
  data: () => ({
    controls: {
      scene: null,
      camera: null,
      renderer: null,
      rotationSpeed: 0.02
    }
  }),
  created () {
    this.$nextTick(() => {
      this.init()
    })
  },
  methods: {
    init () {
      let { initMesh, controls } = this
      //   const gui = new dat.GUI() // gui监测器

      //   gui.add(controls, 'rotationSpeed', 0, 0.5)
      initMesh()
    },
    initMesh () {
      this.scene = new Three.Scene() // 场景
      this.camera = new Three.PerspectiveCamera(45, 800 / 600, 0.1, 1000) // 相机.视场，长宽比，近面，远面
      this.camera.position.x = -20
      this.camera.position.y = 40
      this.camera.position.z = 30
      this.camera.lookAt(this.scene.position)

      this.renderer = new Three.WebGLRenderer({ antialias: true })// 渲染器
      this.renderer.setSize(800, 600 - 100)
      this.renderer.shadowMapEnabled = true // 开启阴影

      let axes = new Three.AxisHelper(20) // 坐标轴

      let planeGeometry = new Three.PlaneGeometry(60, 20, 10, 10) // 生成平面
      let planeMaterial = new Three.MeshLambertMaterial({ color: 0xffffff }) // 材质
      let plane = new Three.Mesh(planeGeometry, planeMaterial)
      plane.rotation.x = -0.5 * Math.PI
      plane.position.x = 0
      plane.position.y = 0
      plane.position.z = 0
      plane.receiveShadow = true

      let cubeGeometry = new Three.CubeGeometry(10, 10, 10)
      let cubeMaterial = new Three.MeshLambertMaterial({ color: 0xff0000 })
      this.cube = new Three.Mesh(cubeGeometry, cubeMaterial)
      this.cube.position.x = -4
      this.cube.position.y = 3
      this.cube.position.z = 0
      this.cube.castShadow = true

      let spotLight = new Three.SpotLight(0xffffff)
      spotLight.position.set(-40, 60, -10)
      spotLight.castShadow = true

      this.scene.add(axes) // 场景添加坐标轴
      this.scene.add(plane) // 向该场景中添加物体
      this.scene.add(this.cube)
      this.scene.add(spotLight)

      this.$refs.demo1.append(this.renderer.domElement)
      this.renderScene()
    },
    renderScene () {
      let { controls, cube, scene, camera, renderer } = this
      cube.rotation.x += controls.rotationSpeed
      cube.rotation.y += controls.rotationSpeed
      cube.rotation.z += controls.rotationSpeed
      requestAnimationFrame(this.renderScene)
      renderer.render(scene, camera)
    }
  }
}
</script>